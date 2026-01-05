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

var testAccProtoV6ProviderFactoriesOpenstackInstance = map[string]func() (tfprotov6.ProviderServer, error){
	"waldur": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func TestAccOpenstackInstanceResource_Basic(t *testing.T) {
	uuid := "550e8400-e29b-41d4-a716-446655440000"
	
	// Initialize state variables to track resource state across requests
	availability_zone := "test-availability_zone"
	connect_directly_to_external_network := true
	data_volume_size := 123
	data_volume_type := "test-data_volume_type"
	data_volumes := []map[string]interface{}{
		{
			"url": "http://example.com/api/security-groups/uuid/",
			"name": "test-sg",
		},
	}
	description := "test-description"
	flavor := "test-flavor"
	floating_ips := []map[string]interface{}{
		{
			"url": "http://example.com/api/security-groups/uuid/",
			"name": "test-sg",
		},
	}
	image := "test-image"
	name := "test-name"
	ports := []map[string]interface{}{
		{
			"url": "http://example.com/api/security-groups/uuid/",
			"name": "test-sg",
		},
	}
	security_groups := []map[string]interface{}{
		{
			"url": "http://example.com/api/security-groups/uuid/",
			"name": "test-sg",
		},
	}
	ssh_public_key := "test-ssh_public_key"
	system_volume_size := 123
	system_volume_type := "test-system_volume_type"
	user_data := "test-user_data"
	project := "test-project"
	offering := "test-offering"

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request: %s %s", r.Method, r.URL.Path)
		
		retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", uuid, 1)
		
		switch {
		case r.Method == "POST" && r.URL.Path == "<no value>":
			// Create operation
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"availability_zone": "%s",
	"connect_directly_to_external_network": %t,
	"data_volume_size": %d,
	"data_volume_type": "%s",
	"description": "%s",
	"flavor": "%s",
	"image": "%s",
	"name": "%s",
	"ssh_public_key": "%s",
	"system_volume_size": %d,
	"system_volume_type": "%s",
	"user_data": "%s",
	"project": "%s",
	"offering": "%s",
	"access_url": "dummy-access_url",
	"action": "dummy-action",
	"availability_zone_name": "dummy-availability_zone_name",
	"backend_id": "dummy-backend_id",
	"cores": 123,
	"created": "dummy-created",
	"customer": "dummy-customer",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_uuid": "dummy-customer_uuid",
	"disk": 123,
	"error_message": "dummy-error_message",
	"error_traceback": "dummy-error_traceback",
	"external_address": ["tag1", "tag2"],
	"external_ips": ["tag1", "tag2"],
	"flavor_disk": 123,
	"flavor_name": "dummy-flavor_name",
	"hypervisor_hostname": "dummy-hypervisor_hostname",
	"image_name": "dummy-image_name",
	"internal_ips": ["tag1", "tag2"],
	"is_limit_based": true,
	"is_usage_based": true,
	"key_fingerprint": "dummy-key_fingerprint",
	"key_name": "dummy-key_name",
	"marketplace_category_name": "dummy-marketplace_category_name",
	"marketplace_category_uuid": "dummy-marketplace_category_uuid",
	"marketplace_offering_name": "dummy-marketplace_offering_name",
	"marketplace_offering_uuid": "dummy-marketplace_offering_uuid",
	"marketplace_plan_uuid": "dummy-marketplace_plan_uuid",
	"marketplace_resource_state": "dummy-marketplace_resource_state",
	"marketplace_resource_uuid": "dummy-marketplace_resource_uuid",
	"min_disk": 123,
	"min_ram": 123,
	"modified": "dummy-modified",
	"project_name": "dummy-project_name",
	"project_uuid": "dummy-project_uuid",
	"ram": 123,
	"resource_type": "dummy-resource_type",
	"runtime_state": "dummy-runtime_state",
	"service_name": "dummy-service_name",
	"service_settings": "dummy-service_settings",
	"service_settings_error_message": "dummy-service_settings_error_message",
	"service_settings_state": "dummy-service_settings_state",
	"service_settings_uuid": "dummy-service_settings_uuid",
	"start_time": "dummy-start_time",
	"tenant": "dummy-tenant",
	"tenant_uuid": "dummy-tenant_uuid",
	"url": "dummy-url",
}`, 
			uuid,
			availability_zone,
			connect_directly_to_external_network,
			data_volume_size,
			data_volume_type,
			data_volumes,
			description,
			flavor,
			floating_ips,
			image,
			name,
			ports,
			security_groups,
			ssh_public_key,
			system_volume_size,
			system_volume_type,
			user_data,
			project,
			offering,
			)))
			
		case r.Method == "GET" && r.URL.Path == retrievePath:
			// Read operation - returns current state
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"availability_zone": "%s",
	"connect_directly_to_external_network": %t,
	"data_volume_size": %d,
	"data_volume_type": "%s",
	"description": "%s",
	"flavor": "%s",
	"image": "%s",
	"name": "%s",
	"ssh_public_key": "%s",
	"system_volume_size": %d,
	"system_volume_type": "%s",
	"user_data": "%s",
	"project": "%s",
	"offering": "%s",
	"access_url": "dummy-access_url",
	"action": "dummy-action",
	"availability_zone_name": "dummy-availability_zone_name",
	"backend_id": "dummy-backend_id",
	"cores": 123,
	"created": "dummy-created",
	"customer": "dummy-customer",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_uuid": "dummy-customer_uuid",
	"disk": 123,
	"error_message": "dummy-error_message",
	"error_traceback": "dummy-error_traceback",
	"external_address": ["tag1", "tag2"],
	"external_ips": ["tag1", "tag2"],
	"flavor_disk": 123,
	"flavor_name": "dummy-flavor_name",
	"hypervisor_hostname": "dummy-hypervisor_hostname",
	"image_name": "dummy-image_name",
	"internal_ips": ["tag1", "tag2"],
	"is_limit_based": true,
	"is_usage_based": true,
	"key_fingerprint": "dummy-key_fingerprint",
	"key_name": "dummy-key_name",
	"marketplace_category_name": "dummy-marketplace_category_name",
	"marketplace_category_uuid": "dummy-marketplace_category_uuid",
	"marketplace_offering_name": "dummy-marketplace_offering_name",
	"marketplace_offering_uuid": "dummy-marketplace_offering_uuid",
	"marketplace_plan_uuid": "dummy-marketplace_plan_uuid",
	"marketplace_resource_state": "dummy-marketplace_resource_state",
	"marketplace_resource_uuid": "dummy-marketplace_resource_uuid",
	"min_disk": 123,
	"min_ram": 123,
	"modified": "dummy-modified",
	"project_name": "dummy-project_name",
	"project_uuid": "dummy-project_uuid",
	"ram": 123,
	"resource_type": "dummy-resource_type",
	"runtime_state": "dummy-runtime_state",
	"service_name": "dummy-service_name",
	"service_settings": "dummy-service_settings",
	"service_settings_error_message": "dummy-service_settings_error_message",
	"service_settings_state": "dummy-service_settings_state",
	"service_settings_uuid": "dummy-service_settings_uuid",
	"start_time": "dummy-start_time",
	"tenant": "dummy-tenant",
	"tenant_uuid": "dummy-tenant_uuid",
	"url": "dummy-url",
}`, 
			uuid,
			availability_zone,
			connect_directly_to_external_network,
			data_volume_size,
			data_volume_type,
			data_volumes,
			description,
			flavor,
			floating_ips,
			image,
			name,
			ports,
			security_groups,
			ssh_public_key,
			system_volume_size,
			system_volume_type,
			user_data,
			project,
			offering,
			)))
			
		case r.Method == "PATCH" && r.URL.Path == retrievePath:
			// Update operation
			// Update state variables
			description = "updated-description"
			name = "updated-name"

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`
{
	"uuid": "%s",
	"availability_zone": "%s",
	"connect_directly_to_external_network": %t,
	"data_volume_size": %d,
	"data_volume_type": "%s",
	"description": "%s",
	"flavor": "%s",
	"image": "%s",
	"name": "%s",
	"ssh_public_key": "%s",
	"system_volume_size": %d,
	"system_volume_type": "%s",
	"user_data": "%s",
	"project": "%s",
	"offering": "%s",
	"access_url": "dummy-access_url",
	"action": "dummy-action",
	"availability_zone_name": "dummy-availability_zone_name",
	"backend_id": "dummy-backend_id",
	"cores": 123,
	"created": "dummy-created",
	"customer": "dummy-customer",
	"customer_abbreviation": "dummy-customer_abbreviation",
	"customer_name": "dummy-customer_name",
	"customer_native_name": "dummy-customer_native_name",
	"customer_uuid": "dummy-customer_uuid",
	"disk": 123,
	"error_message": "dummy-error_message",
	"error_traceback": "dummy-error_traceback",
	"external_address": ["tag1", "tag2"],
	"external_ips": ["tag1", "tag2"],
	"flavor_disk": 123,
	"flavor_name": "dummy-flavor_name",
	"hypervisor_hostname": "dummy-hypervisor_hostname",
	"image_name": "dummy-image_name",
	"internal_ips": ["tag1", "tag2"],
	"is_limit_based": true,
	"is_usage_based": true,
	"key_fingerprint": "dummy-key_fingerprint",
	"key_name": "dummy-key_name",
	"marketplace_category_name": "dummy-marketplace_category_name",
	"marketplace_category_uuid": "dummy-marketplace_category_uuid",
	"marketplace_offering_name": "dummy-marketplace_offering_name",
	"marketplace_offering_uuid": "dummy-marketplace_offering_uuid",
	"marketplace_plan_uuid": "dummy-marketplace_plan_uuid",
	"marketplace_resource_state": "dummy-marketplace_resource_state",
	"marketplace_resource_uuid": "dummy-marketplace_resource_uuid",
	"min_disk": 123,
	"min_ram": 123,
	"modified": "dummy-modified",
	"project_name": "dummy-project_name",
	"project_uuid": "dummy-project_uuid",
	"ram": 123,
	"resource_type": "dummy-resource_type",
	"runtime_state": "dummy-runtime_state",
	"service_name": "dummy-service_name",
	"service_settings": "dummy-service_settings",
	"service_settings_error_message": "dummy-service_settings_error_message",
	"service_settings_state": "dummy-service_settings_state",
	"service_settings_uuid": "dummy-service_settings_uuid",
	"start_time": "dummy-start_time",
	"tenant": "dummy-tenant",
	"tenant_uuid": "dummy-tenant_uuid",
	"url": "dummy-url",
}`, 
			uuid,
			availability_zone,
			connect_directly_to_external_network,
			data_volume_size,
			data_volume_type,
			data_volumes,
			description,
			flavor,
			floating_ips,
			image,
			name,
			ports,
			security_groups,
			ssh_public_key,
			system_volume_size,
			system_volume_type,
			user_data,
			project,
			offering,
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
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactoriesOpenstackInstance,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccOpenstackInstanceResourceConfig_Create(server.URL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "id"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "availability_zone", "test-availability_zone"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "connect_directly_to_external_network", "true"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "data_volume_size", "123"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "data_volume_type", "test-data_volume_type"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "description", "test-description"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "flavor", "test-flavor"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "image", "test-image"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "name", "test-name"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "ssh_public_key", "test-ssh_public_key"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "system_volume_size", "123"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "system_volume_type", "test-system_volume_type"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "user_data", "test-user_data"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "project", "test-project"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "offering", "test-offering"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "access_url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "action"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "availability_zone_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "backend_id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "cores"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "created"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_abbreviation"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_native_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "error_message"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "error_traceback"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.1", "tag2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.1", "tag2"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "flavor_disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "flavor_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "hypervisor_hostname"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "image_name"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.1", "tag2"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "is_limit_based"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "is_usage_based"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "key_fingerprint"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "key_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "latitude"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "longitude"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_category_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_category_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_offering_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_offering_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_plan_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_resource_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_resource_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "min_disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "min_ram"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "modified"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "project_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "project_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "ram"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "resource_type"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "runtime_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "server_group"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_error_message"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "start_time"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "tenant"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "tenant_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "volumes"),
				),
			},
			// Update testing
			{
				Config: testAccOpenstackInstanceResourceConfig_Update(server.URL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "id"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "availability_zone", "updated-availability_zone"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "connect_directly_to_external_network", "false"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "data_volume_size", "456"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "data_volume_type", "updated-data_volume_type"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "description", "updated-description"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "flavor", "updated-flavor"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "image", "updated-image"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "name", "updated-name"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "ssh_public_key", "updated-ssh_public_key"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "system_volume_size", "456"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "system_volume_type", "updated-system_volume_type"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "user_data", "updated-user_data"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "project", "updated-project"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "offering", "updated-offering"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "access_url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "action"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "availability_zone_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "backend_id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "cores"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "created"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_abbreviation"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_native_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "customer_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "error_message"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "error_traceback"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_address.1", "tag2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "external_ips.1", "tag2"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "flavor_disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "flavor_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "hypervisor_hostname"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "image_name"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.0", "tag1"),
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "internal_ips.1", "tag2"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "is_limit_based"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "is_usage_based"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "key_fingerprint"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "key_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "latitude"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "longitude"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_category_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_category_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_offering_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_offering_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_plan_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_resource_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "marketplace_resource_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "min_disk"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "min_ram"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "modified"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "project_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "project_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "ram"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "resource_type"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "runtime_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "server_group"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_name"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_error_message"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_state"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "service_settings_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "start_time"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "tenant"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "tenant_uuid"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "volumes"),
				),
			},
			// Delete testing automatically happens via Terraform
		},
	})
}

// testAccOpenstackInstanceResourceConfig_Create returns a Terraform config for testing create
func testAccOpenstackInstanceResourceConfig_Create(endpoint string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

resource "waldur_openstack_instance" "test" {
  availability_zone = "test-availability_zone"
  connect_directly_to_external_network = true
  data_volume_size = 123
  data_volume_type = "test-data_volume_type"
  data_volumes = [{
    url = "http://example.com/api/security-groups/uuid/"
    name = "test-sg"
  }]
  description = "test-description"
  flavor = "test-flavor"
  floating_ips = [{
    url = "http://example.com/api/security-groups/uuid/"
    name = "test-sg"
  }]
  image = "test-image"
  name = "test-name"
  ports = [{
    url = "http://example.com/api/security-groups/uuid/"
    name = "test-sg"
  }]
  security_groups = [{
    url = "http://example.com/api/security-groups/uuid/"
    name = "test-sg"
  }]
  ssh_public_key = "test-ssh_public_key"
  system_volume_size = 123
  system_volume_type = "test-system_volume_type"
  user_data = "test-user_data"
}
`, endpoint)
}

// testAccOpenstackInstanceResourceConfig_Update returns a Terraform config for testing update
func testAccOpenstackInstanceResourceConfig_Update(endpoint string) string {
	return fmt.Sprintf(`
provider "waldur" {
  endpoint = %[1]q
  token    = "test-token"
}

resource "waldur_openstack_instance" "test" {
  description = "updated-description"
  name = "updated-name"
}
`, endpoint)
}
