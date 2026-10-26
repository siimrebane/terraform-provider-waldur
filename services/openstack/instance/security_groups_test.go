package instance

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

const (
	sgSSH     = "https://w/api/openstack-security-groups/aaa/"
	sgWeb     = "https://w/api/openstack-security-groups/bbb/"
	sgDefault = "https://w/api/openstack-security-groups/ddd/"
)

func sgObject(t *testing.T, url, name string, full bool) types.Object {
	t.Helper()
	values := map[string]attr.Value{
		"url":         types.StringValue(url),
		"description": types.StringNull(),
		"name":        types.StringNull(),
		"rules":       types.ListNull(NestedSecurityGroupRuleType()),
		"state":       types.StringNull(),
	}
	if full {
		values["name"] = types.StringValue(name)
		values["description"] = types.StringValue("about " + name)
		values["state"] = types.StringValue("OK")
	}
	o, d := types.ObjectValue(OpenStackSecurityGroupHyperlinkRequestType().AttrTypes, values)
	if d.HasError() {
		t.Fatalf("object: %v", d)
	}
	return o
}

func sgSet(t *testing.T, objs ...types.Object) types.Set {
	t.Helper()
	vals := make([]attr.Value, 0, len(objs))
	for _, o := range objs {
		vals = append(vals, o)
	}
	s, d := types.SetValue(OpenStackSecurityGroupHyperlinkRequestType(), vals)
	if d.HasError() {
		t.Fatalf("set: %v", d)
	}
	return s
}

func urlsOf(t *testing.T, s types.Set) map[string]types.Object {
	t.Helper()
	out := map[string]types.Object{}
	for _, e := range s.Elements() {
		o := e.(types.Object)
		out[objectString(o, "url")] = o
	}
	return out
}

// The API reports ssh, web and the implicit default group with every
// attribute filled; the practitioner wrote {url} for ssh and web.
func TestNormalizeSecurityGroupsURLOnlyDropsDefault(t *testing.T) {
	api := sgSet(t,
		sgObject(t, sgSSH, "ssh", true),
		sgObject(t, sgWeb, "web", true),
		sgObject(t, sgDefault, "default", true),
	)
	wanted := sgSet(t, sgObject(t, sgSSH, "", false), sgObject(t, sgWeb, "", false))

	got := urlsOf(t, normalizeSecurityGroups(context.Background(), api, wanted, false))
	if len(got) != 2 {
		t.Fatalf("want ssh and web, got %v", got)
	}
	for _, u := range []string{sgSSH, sgWeb} {
		o, ok := got[u]
		if !ok {
			t.Fatalf("missing %s", u)
		}
		if !o.Attributes()["name"].IsNull() || !o.Attributes()["description"].IsNull() || !o.Attributes()["state"].IsNull() {
			t.Fatalf("%s: nested attributes must be null when the plan had only url: %v", u, o)
		}
	}
	if !normalizeSecurityGroups(context.Background(), api, wanted, false).Equal(wanted) {
		t.Fatalf("normalised set must equal the planned set")
	}
}

// A plan that names the group keeps the API's name.
func TestNormalizeSecurityGroupsKeepsWantedAttributes(t *testing.T) {
	api := sgSet(t, sgObject(t, sgSSH, "ssh", true))
	wanted := sgSet(t, sgObject(t, sgSSH, "ssh", true))
	got := urlsOf(t, normalizeSecurityGroups(context.Background(), api, wanted, false))
	if objectString(got[sgSSH], "name") != "ssh" {
		t.Fatalf("name should be kept when it was wanted: %v", got[sgSSH])
	}
}

// Read: the prior state may hold full objects from an older provider; the
// result is url-only anyway, and a wanted default group stays.
func TestNormalizeSecurityGroupsReadIsURLOnly(t *testing.T) {
	api := sgSet(t, sgObject(t, sgSSH, "ssh", true), sgObject(t, sgDefault, "default", true))
	prior := sgSet(t, sgObject(t, sgSSH, "ssh", true), sgObject(t, sgDefault, "default", true))
	got := urlsOf(t, normalizeSecurityGroups(context.Background(), api, prior, true))
	if len(got) != 2 {
		t.Fatalf("default was in the prior state and must stay: %v", got)
	}
	for u, o := range got {
		if !o.Attributes()["name"].IsNull() {
			t.Fatalf("%s: Read must store url only", u)
		}
	}
}

func TestNormalizeSecurityGroupsNullPassesThrough(t *testing.T) {
	n := types.SetNull(OpenStackSecurityGroupHyperlinkRequestType())
	if !normalizeSecurityGroups(context.Background(), n, n, true).IsNull() {
		t.Fatal("null in, null out")
	}
}

// The Waldur actions read a JSON object with one list field, and the
// security-group list holds URLs, not objects.
func TestActionPayloadShapes(t *testing.T) {
	u1, u2 := sgSSH, sgWeb
	b, err := json.Marshal(OpenstackInstanceUpdateSecurityGroupsActionRequest{
		SecurityGroups: []common.OpenStackSecurityGroupHyperlinkRequest{{Url: &u1}, {Url: &u2}, {}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"security_groups":["`+u1+`","`+u2+`"]}` {
		t.Fatalf("security_groups payload: %s", b)
	}

	subnet := "https://w/api/openstack-subnets/s/"
	b, _ = json.Marshal(OpenstackInstanceUpdatePortsActionRequest{Ports: []common.OpenStackCreateInstancePortRequest{{Subnet: &subnet}}})
	var ports map[string][]map[string]any
	if err := json.Unmarshal(b, &ports); err != nil || len(ports["ports"]) != 1 || ports["ports"][0]["subnet"] != subnet {
		t.Fatalf("ports payload: %s (%v)", b, err)
	}

	b, _ = json.Marshal(OpenstackInstanceUpdateFloatingIpsActionRequest{})
	if string(b) != `{"floating_ips":[]}` {
		t.Fatalf("floating_ips payload: %s", b)
	}
}

func TestSecurityGroupLinks(t *testing.T) {
	s := sgSet(t, sgObject(t, sgSSH, "", false), sgObject(t, sgWeb, "web", true))
	links := securityGroupLinks(s)
	if len(links) != 2 {
		t.Fatalf("want 2 links, got %d", len(links))
	}
	for _, l := range links {
		if l.Url == nil || l.Name != nil || l.Rules != nil {
			t.Fatalf("links must carry the url only: %+v", l)
		}
	}
	if len(securityGroupLinks(types.SetNull(OpenStackSecurityGroupHyperlinkRequestType()))) != 0 {
		t.Fatal("null set gives an empty list")
	}
}
