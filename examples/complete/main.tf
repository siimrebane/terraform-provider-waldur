# Complete Example
# Demonstrates all data sources and resource options

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

# =============================================================================
# DATA SOURCES
# =============================================================================

# Look up project by name
data "waldur_structure_project" "project" {
  name = var.project_name
}

# Look up tenant (OpenStack project)
data "waldur_openstack_tenant" "tenant" {
  project_uuid = data.waldur_structure_project.project.id
}

# Look up available flavors
data "waldur_openstack_flavor" "small" {
  name        = "m1.small"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_flavor" "medium" {
  name        = "m1.medium"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Look up available images
data "waldur_openstack_image" "ubuntu22" {
  name        = "Ubuntu 22.04"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_image" "ubuntu24" {
  name        = "Ubuntu 24.04 x86_64"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Look up security groups
data "waldur_openstack_security_group" "default" {
  name        = "default"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_security_group" "ssh" {
  name        = "ssh"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_security_group" "web" {
  name        = "web"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# =============================================================================
# RESOURCES
# =============================================================================

# Create web server VM
resource "waldur_openstack_instance" "web" {
  name    = "${var.environment}-web-server"
  project = data.waldur_structure_project.project.id

  offering           = var.offering_uuid
  flavor             = "${var.waldur_endpoint}/api/openstack-flavors/${data.waldur_openstack_flavor.small.id}/"
  image              = "${var.waldur_endpoint}/api/openstack-images/${data.waldur_openstack_image.ubuntu24.id}/"
  system_volume_size = 20480  # 20 GB
  system_volume_type = "${var.waldur_endpoint}/api/openstack-volume-types/${var.volume_type_uuid}/"

  ports = ["${var.waldur_endpoint}/api/openstack-subnets/${var.subnet_uuid}/"]

  security_groups = [
    "${var.waldur_endpoint}/api/openstack-security-groups/${data.waldur_openstack_security_group.default.id}/",
    "${var.waldur_endpoint}/api/openstack-security-groups/${data.waldur_openstack_security_group.ssh.id}/",
    "${var.waldur_endpoint}/api/openstack-security-groups/${data.waldur_openstack_security_group.web.id}/",
  ]

  ssh_public_key = "${var.waldur_endpoint}/api/keys/${var.ssh_key_uuid}/"
}

# Create app server VM with larger flavor
resource "waldur_openstack_instance" "app" {
  name    = "${var.environment}-app-server"
  project = data.waldur_structure_project.project.id

  offering           = var.offering_uuid
  flavor             = "${var.waldur_endpoint}/api/openstack-flavors/${data.waldur_openstack_flavor.medium.id}/"
  image              = "${var.waldur_endpoint}/api/openstack-images/${data.waldur_openstack_image.ubuntu24.id}/"
  system_volume_size = 40960  # 40 GB
  system_volume_type = "${var.waldur_endpoint}/api/openstack-volume-types/${var.volume_type_uuid}/"

  ports = ["${var.waldur_endpoint}/api/openstack-subnets/${var.subnet_uuid}/"]

  security_groups = [
    "${var.waldur_endpoint}/api/openstack-security-groups/${data.waldur_openstack_security_group.default.id}/",
    "${var.waldur_endpoint}/api/openstack-security-groups/${data.waldur_openstack_security_group.ssh.id}/",
  ]

  ssh_public_key = "${var.waldur_endpoint}/api/keys/${var.ssh_key_uuid}/"
}

# =============================================================================
# OUTPUTS
# =============================================================================

output "project_id" {
  description = "Project UUID"
  value       = data.waldur_structure_project.project.id
}

output "tenant_id" {
  description = "Tenant UUID"
  value       = data.waldur_openstack_tenant.tenant.id
}

output "web_server" {
  description = "Web server details"
  value = {
    id         = waldur_openstack_instance.web.id
    name       = waldur_openstack_instance.web.name
    internal_ips = waldur_openstack_instance.web.internal_ips
    state      = waldur_openstack_instance.web.state
  }
}

output "app_server" {
  description = "App server details"
  value = {
    id         = waldur_openstack_instance.app.id
    name       = waldur_openstack_instance.app.name
    internal_ips = waldur_openstack_instance.app.internal_ips
    state      = waldur_openstack_instance.app.state
  }
}

output "security_groups" {
  description = "Available security group UUIDs"
  value = {
    default = data.waldur_openstack_security_group.default.id
    ssh     = data.waldur_openstack_security_group.ssh.id
    web     = data.waldur_openstack_security_group.web.id
  }
}
