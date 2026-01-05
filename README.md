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
  token    = var.waldur_token  # API token from Waldur
}
```

### Authentication

Get your API token from Waldur:
1. Log in to your Waldur instance
2. Go to Profile → API Token
3. Copy the token

**Security:** Never commit tokens to version control. Use environment variables or tfvars files:

```bash
export TF_VAR_waldur_token="your-token-here"
```

## Quick Start

```hcl
# Look up existing project
data "waldur_structure_project" "myproject" {
  name = "My Project Name"
}

# Look up tenant (OpenStack project)
data "waldur_openstack_tenant" "tenant" {
  project_uuid = data.waldur_structure_project.myproject.id
}

# Look up flavor and image
data "waldur_openstack_flavor" "small" {
  name        = "m1.small"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

data "waldur_openstack_image" "ubuntu" {
  name        = "Ubuntu 22.04"
  tenant_uuid = data.waldur_openstack_tenant.tenant.id
}

# Create a VM
resource "waldur_openstack_instance" "myvm" {
  name    = "my-terraform-vm"
  project = data.waldur_structure_project.myproject.id

  offering           = "your-offering-uuid"
  flavor             = "https://your-waldur/api/openstack-flavors/${data.waldur_openstack_flavor.small.id}/"
  image              = "https://your-waldur/api/openstack-images/${data.waldur_openstack_image.ubuntu.id}/"
  system_volume_size = 20480  # 20GB in MiB

  ports = ["https://your-waldur/api/openstack-subnets/subnet-uuid/"]
}

output "vm_ip" {
  value = waldur_openstack_instance.myvm.internal_ips
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

To find the UUIDs needed for your configuration, you can use the Waldur API:

```bash
# List projects
curl -H "Authorization: Token YOUR_TOKEN" \
  "https://your-waldur/api/projects/"

# List offerings (needed for VM creation)
curl -H "Authorization: Token YOUR_TOKEN" \
  "https://your-waldur/api/marketplace-public-offerings/?type=OpenStack.Instance"

# List subnets
curl -H "Authorization: Token YOUR_TOKEN" \
  "https://your-waldur/api/openstack-subnets/?tenant_uuid=TENANT_UUID"
```

## License

MIT License - see [LICENSE](./LICENSE) for details.

## Contributing

Contributions are welcome! Please open an issue or pull request.
