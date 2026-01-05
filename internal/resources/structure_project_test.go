package resources_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/waldur/terraform-waldur-provider/internal/provider"
)

var testAccProtoV6ProviderFactoriesStructureProject = map[string]func() (tfprotov6.ProviderServer, error){
	"waldur": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func TestAccStructureProjectResource_Basic(t *testing.T) {
	uuid := "550e8400-e29b-41d4-a716-446655440000"
	
	// Initialize state variables to track resource state across requests
	backend_id := "test-backend_id"
	customer := "test-customer"
	description := "test-description"
	end_date := "test-end_date"
	grace_period_days := 123
	image := "test-image"
	is_industry := true
	name := "test-name"
	slug := "test-slug"
	staff_notes := "test-staff_notes"
	start_date := "test-start_date"
	type := "test-type"

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request: %s %s", r.Method, r.URL.Path)
		
		retrievePath := strings.Replace("/api/projects/{uuid}/", "{uuid}", uuid, 1)
		
		switch {
		case r.Method == "POST" && r.URL.Path == "/api/projects/":
			// Create operation
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"backend_id": "%s",
	"customer": "%s",
	"description": "%s",
	"end_date": "%s",
	"grace_period_days": %d,
	"image": "%s",
	"is_industry": %t,
	"name": "%s",
	"slug": "%s",
	"staff_notes": "%s",
	"start_date": "%s",
	"type": "%s",
	"created": "dummy-created",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_display_billing_info_in_projects": true,
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_slug": "dummy-customer_slug",
	"customer_uuid": "dummy-customer_uuid",
	"end_date_requested_by": "dummy-end_date_requested_by",
	"is_removed": true,
	"max_service_accounts": 123,
	"oecd_fos_2007_label": "dummy-oecd_fos_2007_label",
	"resources_count": 123,
	"type_name": "dummy-type_name",
	"type_uuid": "dummy-type_uuid",
	"url": "dummy-url"
}`, 
			uuid,
			backend_id,
			customer,
			description,
			end_date,
			grace_period_days,
			image,
			is_industry,
			name,
			slug,
			staff_notes,
			start_date,
			type,
			)))
			
		case r.Method == "GET" && r.URL.Path == retrievePath:
			// Read operation - returns current state
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"backend_id": "%s",
	"customer": "%s",
	"description": "%s",
	"end_date": "%s",
	"grace_period_days": %d,
	"image": "%s",
	"is_industry": %t,
	"name": "%s",
	"slug": "%s",
	"staff_notes": "%s",
	"start_date": "%s",
	"type": "%s",
	"created": "dummy-created",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_display_billing_info_in_projects": true,
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_slug": "dummy-customer_slug",
	"customer_uuid": "dummy-customer_uuid",
	"end_date_requested_by": "dummy-end_date_requested_by",
	"is_removed": true,
	"max_service_accounts": 123,
	"oecd_fos_2007_label": "dummy-oecd_fos_2007_label",
	"resources_count": 123,
	"type_name": "dummy-type_name",
	"type_uuid": "dummy-type_uuid",
	"url": "dummy-url"
}`, 
			uuid,
			backend_id,
			customer,
			description,
			end_date,
			grace_period_days,
			image,
			is_industry,
			name,
			slug,
			staff_notes,
			start_date,
			type,
			)))
			
		case r.Method == "PATCH" && r.URL.Path == retrievePath:
			// Update operation
			// Update state variables
			backend_id = "updated-backend_id"
			customer = "updated-customer"
			description = "updated-description"
			end_date = "updated-end_date"
			grace_period_days = 456
			image = "updated-image"
			is_industry = false
			name = "updated-name"
			slug = "updated-slug"
			staff_notes = "updated-staff_notes"
			start_date = "updated-start_date"
			type = "updated-type"

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"backend_id": "%s",
	"customer": "%s",
	"description": "%s",
	"end_date": "%s",
	"grace_period_days": %d,
	"image": "%s",
	"is_industry": %t,
	"name": "%s",
	"slug": "%s",
	"staff_notes": "%s",
	"start_date": "%s",
	"type": "%s",
	"created": "dummy-created",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_display_billing_info_in_projects": true,
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_slug": "dummy-customer_slug",
	"customer_uuid": "dummy-customer_uuid",
	"end_date_requested_by": "dummy-end_date_requested_by",
	"is_removed": true,
	"max_service_accounts": 123,
	"oecd_fos_2007_label": "dummy-oecd_fos_2007_label",
	"resources_count": 123,
	"type_name": "dummy-type_name",
	"type_uuid": "dummy-type_uuid",
	"url": "dummy-url"
}`, 
			uuid,
			backend_id,
			customer,
			description,
			end_date,
			grace_period_days,
			image,
			is_industry,
			name,
			slug,
			staff_notes,
			start_date,
			type,
			)))
			
		case r.Method == "DELETE" && r.URL.Path == retrievePath:
			// Delete operation
			w.WriteHeader(http.StatusNoContent)
			
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactoriesStructureProject,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccStructureProjectResourceConfig_Create(server.URL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "id"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "backend_id", "test-backend_id"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "customer", "test-customer"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "description", "test-description"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "end_date", "test-end_date"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "grace_period_days", "123"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "image", "test-image"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "is_industry", "true"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "name", "test-name"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "slug", "test-slug"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "staff_notes", "test-staff_notes"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "start_date", "test-start_date"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "type", "test-type"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "created"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_abbreviation"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_display_billing_info_in_projects"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_native_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_slug"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_uuid"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "end_date_requested_by"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "is_removed"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "max_service_accounts"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "oecd_fos_2007_label"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "project_credit"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "resources_count"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "type_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "type_uuid"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "url"),
				),
			},
			// Update testing
			{
				Config: testAccStructureProjectResourceConfig_Update(server.URL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "id"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "backend_id", "updated-backend_id"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "customer", "updated-customer"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "description", "updated-description"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "end_date", "updated-end_date"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "grace_period_days", "456"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "image", "updated-image"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "is_industry", "false"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "name", "updated-name"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "slug", "updated-slug"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "staff_notes", "updated-staff_notes"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "start_date", "updated-start_date"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "type", "updated-type"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "created"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_abbreviation"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_display_billing_info_in_projects"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_native_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_slug"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer_uuid"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "end_date_requested_by"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "is_removed"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "max_service_accounts"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "oecd_fos_2007_label"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "project_credit"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "resources_count"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "type_name"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "type_uuid"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "url"),
				),
			},
			// Delete testing automatically happens via Terraform
		},
	})
}

// testAccStructureProjectResourceConfig_Create returns a Terraform config for testing create
func testAccStructureProjectResourceConfig_Create(endpoint string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

resource "waldur_structure_project" "test" {
  backend_id = "test-backend_id"
  customer = "test-customer"
  description = "test-description"
  end_date = "test-end_date"
  grace_period_days = 123
  image = "test-image"
  is_industry = true
  name = "test-name"
  slug = "test-slug"
  staff_notes = "test-staff_notes"
  start_date = "test-start_date"
  type = "test-type"
}
`, endpoint)
}

// testAccStructureProjectResourceConfig_Update returns a Terraform config for testing update
func testAccStructureProjectResourceConfig_Update(endpoint string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

resource "waldur_structure_project" "test" {
  backend_id = "updated-backend_id"
  customer = "updated-customer"
  description = "updated-description"
  end_date = "updated-end_date"
  grace_period_days = 456
  image = "updated-image"
  is_industry = false
  name = "updated-name"
  slug = "updated-slug"
  staff_notes = "updated-staff_notes"
  start_date = "updated-start_date"
  type = "updated-type"
}
`, endpoint)
}
