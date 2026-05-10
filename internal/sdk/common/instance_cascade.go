package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// CascadeInstanceSecurityGroupsToPorts pushes the user's HCL
// `security_groups` set onto every Neutron port the instance owns.
//
// # Why this exists
//
// The Waldur marketplace order's `security_groups` field updates only the
// Waldur instance record. Neutron — which actually enforces the rules —
// stores security groups on the port. Without an explicit per-port
// `update_security_groups` call, freshly-created instances come up with
// empty port-level security groups, which means the public IP is
// unreachable on every port (SSH, ICMP, HTTP) regardless of what the
// resource view in the UI shows.
//
// # What it does
//
//  1. GETs the freshly-created instance to discover its ports.
//  2. For each port, POSTs to
//     /api/openstack-ports/{port-uuid}/update_security_groups/ with the
//     list of security-group URLs from the user's plan, plus the
//     "default" security group for the tenant (Waldur attaches the
//     default group at the instance level implicitly, and the port
//     won't be reachable from the rest of the tenant without it).
//
// The function tolerates the absence of the default group on the
// tenant: if it cannot be resolved, the cascade still proceeds with
// the explicit groups.
func CascadeInstanceSecurityGroupsToPorts(
	ctx context.Context,
	c *client.Client,
	instanceUUID string,
	plannedSecurityGroups types.Set,
) error {
	// Extract the security-group URLs from the planned set.
	urls, err := extractSecurityGroupURLs(ctx, plannedSecurityGroups)
	if err != nil {
		return fmt.Errorf("read planned security_groups: %w", err)
	}
	if len(urls) == 0 {
		return nil
	}

	// Fetch instance to read its ports and tenant.
	// Note: the instance.ports response from Waldur exposes only `url` per
	// port (not `uuid`). We extract the UUID from the URL ourselves.
	var instance struct {
		Tenant *string `json:"tenant"`
		Ports  []struct {
			URL *string `json:"url"`
		} `json:"ports"`
	}
	if err := c.GetURL(ctx, fmt.Sprintf("/api/openstack-instances/%s/", instanceUUID), &instance); err != nil {
		return fmt.Errorf("fetch instance ports: %w", err)
	}
	if len(instance.Ports) == 0 {
		// No ports yet — nothing to cascade onto. This can happen if the
		// order has somehow not provisioned a port; let the next refresh
		// surface that.
		return nil
	}

	// Resolve the tenant's "default" security group and append it to the
	// port-level set only. We deliberately do NOT touch the instance
	// record here — the Waldur "Resource" tab should reflect exactly what
	// the user wrote in HCL, not what the platform adds underneath.
	// Adding default to the instance record after Create would make the
	// returned set's length differ from the plan and Terraform would
	// reject the apply with a "provider produced inconsistent result"
	// error.
	defaultURL := resolveDefaultSecurityGroupURL(ctx, c, instance.Tenant)
	allURLs := urls
	if defaultURL != "" && !containsString(allURLs, defaultURL) {
		allURLs = append(allURLs, defaultURL)
	}

	// Cascade to every port (this is what Neutron actually enforces).
	for _, p := range instance.Ports {
		if p.URL == nil {
			continue
		}
		portUUID := ExtractUUIDFromURL(*p.URL)
		if portUUID == "" {
			continue
		}
		path := fmt.Sprintf("/api/openstack-ports/%s/update_security_groups/", portUUID)
		body := map[string]interface{}{"security_groups": allURLs}
		if err := c.Post(ctx, path, body, nil); err != nil {
			return fmt.Errorf("cascade to port %s: %w", portUUID, err)
		}
	}
	return nil
}

// extractSecurityGroupURLs returns the `url` value from each element of
// the planned security_groups set.
//
// Implementation note: we avoid types.Set.ElementsAs into a struct
// because the planned set's nested attributes have computed/unknown
// fields (description, rules, state) that the decode treats as
// undecodable and silently produces an empty slice on. Working at the
// attr.Value level via Elements() and asserting types.Object lets us
// pluck just the `url` attribute reliably.
func extractSecurityGroupURLs(_ context.Context, set types.Set) ([]string, error) {
	if set.IsNull() || set.IsUnknown() {
		return nil, nil
	}

	var out []string
	for _, elem := range set.Elements() {
		obj, ok := elem.(types.Object)
		if !ok {
			continue
		}
		attrs := obj.Attributes()
		urlAttr, hasURL := attrs["url"]
		if !hasURL {
			continue
		}
		urlStr, ok := urlAttr.(types.String)
		if !ok {
			continue
		}
		if urlStr.IsNull() || urlStr.IsUnknown() {
			continue
		}
		v := urlStr.ValueString()
		if v != "" {
			out = append(out, v)
		}
	}
	return out, nil
}

// resolveDefaultSecurityGroupURL fetches the tenant's "default" security
// group URL. Returns "" if it cannot be resolved (which is fine — the
// cascade will simply not include it).
func resolveDefaultSecurityGroupURL(ctx context.Context, c *client.Client, tenantURL *string) string {
	if tenantURL == nil {
		return ""
	}
	tenantUUID := ExtractUUIDFromURL(*tenantURL)
	if tenantUUID == "" {
		return ""
	}

	var groups []struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	}
	path := fmt.Sprintf("/api/openstack-security-groups/?tenant_uuid=%s&name_exact=default", tenantUUID)
	if err := c.GetURL(ctx, path, &groups); err != nil {
		return ""
	}
	for _, g := range groups {
		if g.Name != nil && *g.Name == "default" && g.URL != nil {
			return *g.URL
		}
	}
	return ""
}

func containsString(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
