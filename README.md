# Terraform Provider for Waldur

Terraform provider for managing resources in [Waldur](https://waldur.com/) - a hybrid cloud management platform. This provider allows you to create and manage OpenStack virtual machines through Waldur's marketplace.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- A Waldur account with API access

## Installation

```hcl
terraform {
  required_providers {
    waldur = {
      source  = "siimrebane/waldur"
      version = "~> 0.1"
    }
  }
}
```

Then run `terraform init` to download the provider.

## Configuration

```hcl
provider "waldur" {
  endpoint = "https://your-waldur-instance.com"
  token    = var.waldur_token
}

variable "waldur_token" {
  description = "Waldur API token"
  type        = string
  sensitive   = true
}
```

### Authentication

Get your API token from Waldur:
1. Log in to your Waldur instance
2. Go to Profile → API Token
3. Copy the token

**Security:** Never commit tokens to version control. Use environment variables:

```bash
export TF_VAR_waldur_token="your-token-here"
```

## Quick Start - Create a VM

```hcl
terraform {
  required_providers {
    waldur = {
      source  = "siimrebane/waldur"
      version = "~> 0.1"
    }
  }
}

provider "waldur" {
  endpoint = "https://your-waldur-instance.com"
  token    = var.waldur_token
}

variable "waldur_token" {
  type      = string
  sensitive = true
}

# Look up your project
data "waldur_structure_project" "project" {
  name = "My Project Name"
}

# Look up tenant (OpenStack project)
data "waldur_openstack_tenant" "tenant" {
  project_uuid = data.waldur_structure_project.project.id
}

# Look up flavor and image
data "waldur_openstack_flavor" "flavor" {
  name        = "m1.small"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_image" "image" {
  name        = "Ubuntu 22.04"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Create a VM
resource "waldur_openstack_instance" "vm" {
  name    = "my-terraform-vm"
  project = data.waldur_structure_project.project.id

  offering           = "your-offering-uuid"
  flavor             = "${var.waldur_endpoint}/api/openstack-flavors/${data.waldur_openstack_flavor.flavor.id}/"
  image              = "${var.waldur_endpoint}/api/openstack-images/${data.waldur_openstack_image.image.id}/"
  system_volume_size = 20480  # 20GB in MiB
  system_volume_type = "${var.waldur_endpoint}/api/openstack-volume-types/your-volume-type-uuid/"

  ports          = ["${var.waldur_endpoint}/api/openstack-subnets/your-subnet-uuid/"]
  ssh_public_key = "${var.waldur_endpoint}/api/keys/your-ssh-key-uuid/"
}

output "vm_id" {
  value = waldur_openstack_instance.vm.id
}

output "vm_ip" {
  value = waldur_openstack_instance.vm.internal_ips
}
```

Then run:
```bash
export TF_VAR_waldur_token="your-token"
terraform init
terraform plan
terraform apply
```

## ETAIS Example (minu.etais.ee)

Working example for Estonian Scientific Computing Infrastructure:

```hcl
terraform {
  required_providers {
    waldur = {
      source  = "siimrebane/waldur"
      version = "~> 0.1"
    }
  }
}

provider "waldur" {
  endpoint = "https://minu.etais.ee"
  token    = var.waldur_token
}

variable "waldur_token" {
  type      = string
  sensitive = true
}

# Look up your project by name
data "waldur_structure_project" "project" {
  name = "Your Project Name"
}

# Look up tenant
data "waldur_openstack_tenant" "tenant" {
  project_uuid = data.waldur_structure_project.project.id
}

# Look up flavor and image
data "waldur_openstack_flavor" "flavor" {
  name        = "g1.small1"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_image" "ubuntu" {
  name        = "Ubuntu 24.04 x86_64"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Create VM
resource "waldur_openstack_instance" "vm" {
  name    = "my-etais-vm"
  project = data.waldur_structure_project.project.id

  # Get offering UUID from: https://minu.etais.ee/api/marketplace-public-offerings/?type=OpenStack.Instance
  offering           = "972a1ad09a68446e9e1a1824529b7142"
  flavor             = "https://minu.etais.ee/api/openstack-flavors/${data.waldur_openstack_flavor.flavor.id}/"
  image              = "https://minu.etais.ee/api/openstack-images/${data.waldur_openstack_image.ubuntu.id}/"
  system_volume_size = 10240  # 10GB in MiB
  system_volume_type = "https://minu.etais.ee/api/openstack-volume-types/6e48614ac58d40bba626cec8158f432b/"

  # Get subnet UUID from: https://minu.etais.ee/api/openstack-subnets/?tenant_uuid=YOUR_TENANT
  ports = ["https://minu.etais.ee/api/openstack-subnets/YOUR_SUBNET_UUID/"]

  # Get SSH key UUID from: https://minu.etais.ee/api/keys/
  ssh_public_key = "https://minu.etais.ee/api/keys/YOUR_SSH_KEY_UUID/"
}

output "vm_id" {
  value = waldur_openstack_instance.vm.id
}

output "vm_ip" {
  value = waldur_openstack_instance.vm.internal_ips
}
```

## Data Sources

| Name | Description |
|------|-------------|
| `waldur_structure_project` | Look up a project by name |
| `waldur_openstack_tenant` | Look up an OpenStack tenant |
| `waldur_openstack_flavor` | Look up VM flavors (sizes) |
| `waldur_openstack_image` | Look up OS images |
| `waldur_openstack_instance` | Look up existing VMs |
| `waldur_openstack_security_group` | Look up security groups |
| `waldur_openstack_floating_ip` | Look up floating IPs |

## Resources

| Name | Description |
|------|-------------|
| `waldur_openstack_instance` | Create OpenStack VMs via marketplace |

## Examples

See the [examples](./examples/) folder for complete working examples:

- [basic-vm](./examples/basic-vm/) - Simple VM creation
- [complete](./examples/complete/) - Full example with all data sources

## Finding UUIDs

Use the Waldur API to find required UUIDs:

```bash
# Set your token
export TOKEN="your-api-token"
export WALDUR="https://your-waldur-instance.com"

# List your projects
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/projects/"

# Get tenant UUID (after you have project UUID)
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/openstack-tenants/?project_uuid=PROJECT_UUID"

# List offerings (needed for VM creation)
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/marketplace-public-offerings/?type=OpenStack.Instance"

# List subnets (after you have tenant UUID)
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/openstack-subnets/?tenant_uuid=TENANT_UUID"

# List volume types
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/openstack-volume-types/?tenant_uuid=TENANT_UUID"

# List your SSH keys
curl -H "Authorization: Token $TOKEN" "$WALDUR/api/keys/"
```

## Known Issues

After `terraform apply`, you may see errors about "unknown values" for some fields (like `access_url`, `hypervisor_hostname`). These are cosmetic - the VM is still created successfully. Check with `terraform state show` or in the Waldur UI.

## License

MIT License

## Contributing

Contributions are welcome! Please open an issue or pull request on [GitHub](https://github.com/siimrebane/terraform-provider-waldur).
