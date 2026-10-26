package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// normalizeSecurityGroups shapes the instance's security_groups set the way
// the practitioner wrote it, so that a refresh does not plan an update on
// every run.
//
// The API answers with full security-group objects (name, description,
// rules, state) and always includes the tenant's implicit "default" group.
// HCL normally lists {url = ...} only. A set compares whole elements, so
// storing the API's objects made every plan show the set as changed.
//
// The result keeps, per group the API reports, the url plus those nested
// attributes the wanted element with the same url had set; the rest is
// null. With urlOnly every nested attribute but url is null, whatever the
// wanted element had: Read uses that, because the prior state may still
// hold the full objects an earlier provider version stored. The implicit
// "default" group is left out unless it was wanted. wanted is the prior
// state on Read and the plan on Create and Update.
func normalizeSecurityGroups(ctx context.Context, fromAPI types.Set, wanted types.Set, urlOnly bool) types.Set {
	if fromAPI.IsNull() || fromAPI.IsUnknown() {
		return fromAPI
	}
	wantedByURL := map[string]types.Object{}
	if !wanted.IsNull() && !wanted.IsUnknown() {
		for _, e := range wanted.Elements() {
			if obj, ok := e.(types.Object); ok {
				if u := objectString(obj, "url"); u != "" {
					wantedByURL[u] = obj
				}
			}
		}
	}

	elemType := OpenStackSecurityGroupHyperlinkRequestType()
	out := make([]attr.Value, 0, len(fromAPI.Elements()))
	for _, e := range fromAPI.Elements() {
		obj, ok := e.(types.Object)
		if !ok {
			continue
		}
		u := objectString(obj, "url")
		if u == "" {
			continue
		}
		w, isWanted := wantedByURL[u]
		if objectString(obj, "name") == "default" && !isWanted {
			continue
		}
		values := map[string]attr.Value{
			"url":         types.StringValue(u),
			"description": types.StringNull(),
			"name":        types.StringNull(),
			"rules":       types.ListNull(NestedSecurityGroupRuleType()),
			"state":       types.StringNull(),
		}
		if isWanted && !urlOnly {
			for _, k := range []string{"description", "name", "rules", "state"} {
				wv, ok := w.Attributes()[k]
				if !ok || wv.IsNull() || wv.IsUnknown() {
					continue
				}
				if av, ok := obj.Attributes()[k]; ok && !av.IsNull() {
					values[k] = av
				}
			}
		}
		o, d := types.ObjectValue(elemType.AttrTypes, values)
		if d.HasError() {
			continue
		}
		out = append(out, o)
	}
	set, d := types.SetValue(elemType, out)
	if d.HasError() {
		return fromAPI
	}
	return set
}

// securityGroupLinks turns a security_groups set into the url-only list the
// update_security_groups action is sent. Elements without a known url are
// skipped.
func securityGroupLinks(set types.Set) []common.OpenStackSecurityGroupHyperlinkRequest {
	out := []common.OpenStackSecurityGroupHyperlinkRequest{}
	if set.IsNull() || set.IsUnknown() {
		return out
	}
	for _, e := range set.Elements() {
		if obj, ok := e.(types.Object); ok {
			if u := objectString(obj, "url"); u != "" {
				u := u
				out = append(out, common.OpenStackSecurityGroupHyperlinkRequest{Url: &u})
			}
		}
	}
	return out
}

// objectString returns the known, non-null string attribute k of obj, or "".
func objectString(obj types.Object, k string) string {
	v, ok := obj.Attributes()[k]
	if !ok {
		return ""
	}
	s, ok := v.(types.String)
	if !ok || s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}
