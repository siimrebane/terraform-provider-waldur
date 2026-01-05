# Basic VM Example
# Creates a single OpenStack VM via Waldur marketplace

terraform {
  required_providers {
    waldur = {
      source  = "siimrebane/waldur"
      version = "~> 0.1"
    }
  }
}

provider "waldur" {
  endpoint = var.waldur_endpoint
  token    = var.waldur_token
}

# Look up your project
data "waldur_structure_project" "project" {
  name = var.project_name
}

# Look up tenant (OpenStack project in Waldur)
data "waldur_openstack_tenant" "tenant" {
  project_uuid = data.waldur_structure_project.project.id
}

# Look up flavor and image
data "waldur_openstack_flavor" "flavor" {
  name        = var.flavor_name
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_image" "image" {
  name        = var.image_name
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Create the VM
resource "waldur_openstack_instance" "vm" {
  name    = var.vm_name
  project = data.waldur_structure_project.project.id

  offering           = var.offering_uuid
  flavor             = "${var.waldur_endpoint}/api/openstack-flavors/${data.waldur_openstack_flavor.flavor.id}/"
  image              = "${var.waldur_endpoint}/api/openstack-images/${data.waldur_openstack_image.image.id}/"
  system_volume_size = var.volume_size
  system_volume_type = "${var.waldur_endpoint}/api/openstack-volume-types/${var.volume_type_uuid}/"

  ports = ["${var.waldur_endpoint}/api/openstack-subnets/${var.subnet_uuid}/"]

  ssh_public_key = "${var.waldur_endpoint}/api/keys/${var.ssh_key_uuid}/"
}
