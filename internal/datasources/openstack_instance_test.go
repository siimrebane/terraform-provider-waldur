package datasources_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	
	"github.com/waldur/terraform-waldur-provider/internal/provider"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"waldur": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func TestAccStructureProjectDataSource_ByUUID(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("ByUUID Request: %s %s", r.Method, r.URL.Path)
		
		// Expected path with UUID substituted
		expectedPath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", "550e8400-e29b-41d4-a716-446655440000", 1)
		if !strings.HasSuffix(expectedPath, "/") {
			expectedPath += "/"
		}
		t.Logf("Expected path: %s, got: %s", expectedPath, r.URL.Path)
		
		if r.URL.Path == expectedPath {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"uuid": "550e8400-e29b-41d4-a716-446655440000",
				"name": "Test OpenstackInstance"
			}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing by UUID
			{
				Config: testAccOpenstackInstanceDataSourceConfig_ByUUID(server.URL, "550e8400-e29b-41d4-a716-446655440000"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.waldur_openstack_instance.test", "id", "550e8400-e29b-41d4-a716-446655440000"),
					resource.TestCheckResourceAttr("data.waldur_openstack_instance.test", "name", "Test OpenstackInstance"),
				),
			},
		},
	})
}

func TestAccOpenstackInstanceDataSource_ByName(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request: %s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery)
		
		queryFilter := r.URL.Query().Get("name_exact")
		if queryFilter == "" {
			queryFilter = r.URL.Query().Get("name")
		}

		// Use list path for filtering
		t.Logf("Expected path: %s, got: %s", "/api/openstack-instances/", r.URL.Path)
		t.Logf("Expected filter: %s, got: %s", "Test OpenstackInstance", queryFilter)
		if r.URL.Path == "/api/openstack-instances/" && queryFilter == "Test OpenstackInstance" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{
				"uuid": "550e8400-e29b-41d4-a716-446655440000",
				"name": "Test OpenstackInstance"
			}]`))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[]`))
		}
	}))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing by name
			{
				Config: testAccOpenstackInstanceDataSourceConfig_ByName(server.URL, "Test OpenstackInstance"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.waldur_openstack_instance.test", "id", "550e8400-e29b-41d4-a716-446655440000"),
					resource.TestCheckResourceAttr("data.waldur_openstack_instance.test", "name", "Test OpenstackInstance"),
				),
			},
		},
	})
}

func TestAccOpenstackInstanceDataSource_NotFound(t *testing.T) {
	// Create a test server that returns empty results
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccOpenstackInstanceDataSourceConfig_ByName(server.URL, "Nonexistent"),
				ExpectError: regexp.MustCompile("Instance Not Found"),
			},
		},
	})
}

func TestAccOpenstackInstanceDataSource_MultipleResults(t *testing.T) {
	// Create a test server that returns multiple results
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/openstack-instances/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{"uuid": "uuid-1", "name": "Ambiguous"},
				{"uuid": "uuid-2", "name": "Ambiguous"}
			]`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccOpenstackInstanceDataSourceConfig_ByName(server.URL, "Ambiguous"),
				ExpectError: regexp.MustCompile("Multiple Instances Found"),
			},
		},
	})
}

// testAccOpenstackInstanceDataSourceConfig_ByUUID returns a Terraform config for testing lookup by UUID
func testAccOpenstackInstanceDataSourceConfig_ByUUID(endpoint, uuid string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

data "waldur_openstack_instance" "test" {
  id = %[2]q
}
`, endpoint, uuid)
}

// testAccOpenstackInstanceDataSourceConfig_ByName returns a Terraform config for testing lookup by name
func testAccOpenstackInstanceDataSourceConfig_ByName(endpoint, name string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

data "waldur_openstack_instance" "test" {
  name = %[2]q
}
`, endpoint, name)
}

