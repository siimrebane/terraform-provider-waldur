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
		expectedPath := strings.Replace("/api/openstack-flavors/{uuid}/", "{uuid}", "550e8400-e29b-41d4-a716-446655440000", 1)
		if !strings.HasSuffix(expectedPath, "/") {
			expectedPath += "/"
		}
		t.Logf("Expected path: %s, got: %s", expectedPath, r.URL.Path)
		
		if r.URL.Path == expectedPath {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"uuid": "550e8400-e29b-41d4-a716-446655440000",
				"name": "Test OpenstackFlavor"
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
				Config: testAccOpenstackFlavorDataSourceConfig_ByUUID(server.URL, "550e8400-e29b-41d4-a716-446655440000"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.waldur_openstack_flavor.test", "id", "550e8400-e29b-41d4-a716-446655440000"),
					resource.TestCheckResourceAttr("data.waldur_openstack_flavor.test", "name", "Test OpenstackFlavor"),
				),
			},
		},
	})
}

func TestAccOpenstackFlavorDataSource_ByName(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request: %s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery)
		
		queryFilter := r.URL.Query().Get("name_exact")
		if queryFilter == "" {
			queryFilter = r.URL.Query().Get("name")
		}

		// Use list path for filtering
		t.Logf("Expected path: %s, got: %s", "/api/openstack-flavors/", r.URL.Path)
		t.Logf("Expected filter: %s, got: %s", "Test OpenstackFlavor", queryFilter)
		if r.URL.Path == "/api/openstack-flavors/" && queryFilter == "Test OpenstackFlavor" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{
				"uuid": "550e8400-e29b-41d4-a716-446655440000",
				"name": "Test OpenstackFlavor"
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
				Config: testAccOpenstackFlavorDataSourceConfig_ByName(server.URL, "Test OpenstackFlavor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.waldur_openstack_flavor.test", "id", "550e8400-e29b-41d4-a716-446655440000"),
					resource.TestCheckResourceAttr("data.waldur_openstack_flavor.test", "name", "Test OpenstackFlavor"),
				),
			},
		},
	})
}

func TestAccOpenstackFlavorDataSource_NotFound(t *testing.T) {
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
				Config:      testAccOpenstackFlavorDataSourceConfig_ByName(server.URL, "Nonexistent"),
				ExpectError: regexp.MustCompile("Flavor Not Found"),
			},
		},
	})
}

func TestAccOpenstackFlavorDataSource_MultipleResults(t *testing.T) {
	// Create a test server that returns multiple results
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/openstack-flavors/" {
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
				Config:      testAccOpenstackFlavorDataSourceConfig_ByName(server.URL, "Ambiguous"),
				ExpectError: regexp.MustCompile("Multiple Flavors Found"),
			},
		},
	})
}

// testAccOpenstackFlavorDataSourceConfig_ByUUID returns a Terraform config for testing lookup by UUID
func testAccOpenstackFlavorDataSourceConfig_ByUUID(endpoint, uuid string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

data "waldur_openstack_flavor" "test" {
  id = %[2]q
}
`, endpoint, uuid)
}

// testAccOpenstackFlavorDataSourceConfig_ByName returns a Terraform config for testing lookup by name
func testAccOpenstackFlavorDataSourceConfig_ByName(endpoint, name string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

data "waldur_openstack_flavor" "test" {
  name = %[2]q
}
`, endpoint, name)
}

