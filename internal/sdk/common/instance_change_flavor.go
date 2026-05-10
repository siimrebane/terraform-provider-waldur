package common

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// ChangeInstanceFlavorWithStopStart performs an in-place flavor change on
// an OpenStack instance.
//
// # Why this exists
//
// Waldur exposes a `change_flavor` action on the instance resource:
//
//	POST /api/openstack-instances/{uuid}/change_flavor/   {"flavor": "<url>"}
//
// The action's documentation explicitly notes that the instance must be
// stopped (RuntimeState=SHUTOFF) before the flavor change can succeed.
// Calling the endpoint against an ACTIVE instance returns an error and
// leaves the VM in its previous flavor.
//
// Without this helper, the only way to resize a VM through Terraform is
// to taint and recreate it — which destroys the volume, the floating IP,
// and any data on the instance. That's the opposite of what a student
// expects when they edit `flavor = "g1.medium2"` in their HCL and run
// `terraform apply`.
//
// What it does
//
//  1. Read the instance's current runtime_state.
//  2. If the instance is ACTIVE, call /stop/ and wait for state=OK.
//     (Waldur transitions state OK -> UPDATING -> OK when stopping; at
//     the second OK the runtime_state is SHUTOFF.)
//  3. POST /change_flavor/ with the new flavor URL.
//  4. Wait for state=OK after the change.
//  5. If the instance was ACTIVE before, call /start/ and wait for
//     state=OK so the VM ends in the same lifecycle position it began.
//
// The helper uses raw HTTP via the base client.Client so it doesn't
// depend on the per-resource Go client packages (which would be a cyclic
// import from the SDK common layer).
func ChangeInstanceFlavorWithStopStart(
	ctx context.Context,
	c *client.Client,
	instanceUUID string,
	newFlavorURL string,
	timeout time.Duration,
) error {
	// 0. Snapshot the instance-level security_groups before we touch
	// anything. Waldur's `start` API has a habit of auto-attaching the
	// tenant's `default` group to the instance record on the way back
	// up if the SG list is anything other than what was originally
	// planned, which makes Terraform's plan-vs-apply consistency check
	// fail with "length changed from N to N+1". We restore the
	// snapshot at the end so the instance record matches what it was
	// before the resize.
	preSGURLs, err := snapshotInstanceSecurityGroupURLs(ctx, c, instanceUUID)
	if err != nil {
		return fmt.Errorf("snapshot security_groups before change_flavor: %w", err)
	}

	// 1. Read current state.
	current, err := getInstanceLifecycle(ctx, c, instanceUUID)
	if err != nil {
		return fmt.Errorf("read instance before change_flavor: %w", err)
	}
	wasRunning := current.RuntimeState == "ACTIVE"

	// 2. If running, stop and wait for SHUTOFF.
	if wasRunning {
		if err := c.Post(ctx, fmt.Sprintf("/api/openstack-instances/%s/stop/", instanceUUID), nil, nil); err != nil {
			return fmt.Errorf("stop instance for change_flavor: %w", err)
		}
		if err := waitForInstance(ctx, c, instanceUUID, "SHUTOFF", timeout); err != nil {
			return fmt.Errorf("wait for SHUTOFF before change_flavor: %w", err)
		}
	}

	// 3. POST /change_flavor/.
	body := map[string]interface{}{"flavor": newFlavorURL}
	if err := c.Post(ctx, fmt.Sprintf("/api/openstack-instances/%s/change_flavor/", instanceUUID), body, nil); err != nil {
		return fmt.Errorf("change_flavor action: %w", err)
	}

	// 4. Wait for state=OK after the change. The runtime_state at this
	// point is SHUTOFF (we never started it), so we wait on State only.
	if err := waitForInstance(ctx, c, instanceUUID, "", timeout); err != nil {
		return fmt.Errorf("wait for OK after change_flavor: %w", err)
	}

	// 5. If the instance was running before, start it back up so the
	// student is left with a VM in the same lifecycle position they
	// edited the flavor of.
	if wasRunning {
		if err := c.Post(ctx, fmt.Sprintf("/api/openstack-instances/%s/start/", instanceUUID), nil, nil); err != nil {
			return fmt.Errorf("restart instance after change_flavor: %w", err)
		}
		if err := waitForInstance(ctx, c, instanceUUID, "ACTIVE", timeout); err != nil {
			return fmt.Errorf("wait for ACTIVE after change_flavor: %w", err)
		}
	}

	// 6. Restore the snapshot. If the post-start instance record
	// already matches the snapshot (Waldur didn't drift it), the call
	// is a no-op as far as the visible state is concerned.
	postSGURLs, err := snapshotInstanceSecurityGroupURLs(ctx, c, instanceUUID)
	if err != nil {
		return fmt.Errorf("read security_groups after change_flavor: %w", err)
	}
	if !sameURLSet(preSGURLs, postSGURLs) {
		body := map[string]interface{}{"security_groups": preSGURLs}
		if err := c.Post(ctx, fmt.Sprintf("/api/openstack-instances/%s/update_security_groups/", instanceUUID), body, nil); err != nil {
			return fmt.Errorf("restore security_groups after change_flavor: %w", err)
		}
		if err := waitForInstance(ctx, c, instanceUUID, "", timeout); err != nil {
			return fmt.Errorf("wait for OK after security_groups restore: %w", err)
		}
	}
	return nil
}

// snapshotInstanceSecurityGroupURLs returns the URLs of all security
// groups currently attached to the instance record (instance level, not
// per-port).
func snapshotInstanceSecurityGroupURLs(ctx context.Context, c *client.Client, instanceUUID string) ([]string, error) {
	var raw struct {
		SecurityGroups []struct {
			URL *string `json:"url"`
		} `json:"security_groups"`
	}
	if err := c.GetURL(ctx, fmt.Sprintf("/api/openstack-instances/%s/", instanceUUID), &raw); err != nil {
		return nil, err
	}
	out := make([]string, 0, len(raw.SecurityGroups))
	for _, sg := range raw.SecurityGroups {
		if sg.URL != nil && *sg.URL != "" {
			out = append(out, *sg.URL)
		}
	}
	return out, nil
}

// sameURLSet returns true if the two slices contain the same URLs
// (order-insensitive).
func sameURLSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[string]struct{}, len(a))
	for _, s := range a {
		seen[s] = struct{}{}
	}
	for _, s := range b {
		if _, ok := seen[s]; !ok {
			return false
		}
	}
	return true
}

// instanceLifecycle is a minimal subset of the instance response that this
// helper needs to make decisions. Defined locally to avoid pulling in the
// generated OpenstackInstanceResponse type (which lives in a different
// package and would create an import cycle).
type instanceLifecycle struct {
	State        string
	RuntimeState string
}

func getInstanceLifecycle(ctx context.Context, c *client.Client, instanceUUID string) (instanceLifecycle, error) {
	var raw struct {
		State        *string `json:"state"`
		RuntimeState *string `json:"runtime_state"`
	}
	if err := c.GetURL(ctx, fmt.Sprintf("/api/openstack-instances/%s/", instanceUUID), &raw); err != nil {
		return instanceLifecycle{}, err
	}
	out := instanceLifecycle{}
	if raw.State != nil {
		out.State = *raw.State
	}
	if raw.RuntimeState != nil {
		out.RuntimeState = *raw.RuntimeState
	}
	return out, nil
}

// waitForInstance blocks until the instance reaches state=OK, optionally
// also requiring a particular runtime_state. Pass "" for runtime_state to
// only wait on state.
func waitForInstance(ctx context.Context, c *client.Client, instanceUUID, requiredRuntime string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			"CREATION_SCHEDULED", "CREATING",
			"UPDATE_SCHEDULED", "UPDATING",
			"DELETION_SCHEDULED", "DELETING",
		},
		Target: []string{"OK"},
		Refresh: func() (interface{}, string, error) {
			lc, err := getInstanceLifecycle(ctx, c, instanceUUID)
			if err != nil {
				return nil, "", err
			}
			if lc.State == "ERRED" {
				return lc, "failed", fmt.Errorf("instance entered ERRED state")
			}
			// State is OK but runtime_state hasn't reached the
			// required value yet — keep polling.
			if lc.State == "OK" && requiredRuntime != "" && lc.RuntimeState != requiredRuntime {
				return lc, "UPDATING", nil
			}
			return lc, lc.State, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
