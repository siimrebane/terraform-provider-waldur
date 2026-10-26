# Terraform Provider for Waldur (siimrebane fork)

A Terraform provider for [Waldur](https://waldur.com/), the platform behind
Estonia's academic cloud [ETAIS](https://etais.ee/) (`minu.etais.ee`).

This is a fork of the upstream
[`waldur/terraform-provider-waldur`](https://github.com/waldur/terraform-provider-waldur).
Version 0.3.x is upstream **v8.1.3** plus the two patches below. Earlier fork
versions (0.1.x, 0.2.x) were based on upstream v0.0.8; see the
[changelog](CHANGELOG.md) for what was dropped.

## Installation

```hcl
terraform {
  required_providers {
    waldur = {
      source  = "siimrebane/waldur"
      version = "~> 0.3"
    }
  }
}

provider "waldur" {
  endpoint = "https://minu.etais.ee"
  token    = var.waldur_token
}
```

`terraform init` downloads the provider from the
[Terraform Registry](https://registry.terraform.io/providers/siimrebane/waldur).
Coming from 0.2.x: `terraform init -upgrade`, and drop any
`lifecycle { ignore_changes = [security_groups] }` the 0.2.x workaround needed.

## What this fork changes

Both patches concern `security_groups` on `waldur_openstack_instance`.
Files: `services/openstack/instance/security_groups.go`,
`internal/sdk/common/instance_cascade.go`, and the call sites in
`services/openstack/instance/resource.go`.

### 1. The groups reach the instance's ports

Waldur records the groups you list on the instance, but it does not put
them on the instance's Neutron port, and the port is what enforces traffic.
On ETAIS a port without security groups gets no DHCP lease: the machine
boots without an address and is unreachable. After create, and after an
`update_security_groups` action, this fork pushes the listed groups plus the
tenant's `default` group to every port of the instance.

### 2. A clean plan

The API answers with full security-group objects (name, description, rules,
state) and always includes the tenant's implicit `default` group, while the
configuration lists `{ url = ... }` only. A set attribute compares whole
elements, so upstream stores the API's objects and every refresh plans the
set as changed. This fork stores the groups the way the configuration wrote
them, url only and without the implicit `default` group unless it was
listed, and sends the action the planned urls.

```hcl
resource "waldur_openstack_instance" "vm" {
  # ...
  security_groups = [
    { url = "https://minu.etais.ee/api/openstack-security-groups/<uuid>/" },  # ssh
    { url = "https://minu.etais.ee/api/openstack-security-groups/<uuid>/" },  # ping
  ]
}
```

Verified against ETAIS (2026-10-26): create, SSH in, a second plan with no
changes, adding and removing a group through `update_security_groups`, a
clean plan again, destroy. A control build without patch 1 produced a
machine with no groups on its port and no DHCP lease.

### Dropped from the 0.2.x fork

- `change_flavor` in place (stop, change flavour, start): not in upstream,
  not kept. A flavour change follows upstream's behaviour.
- The `future_prices` string-or-number fix: upstream now types the field as
  strings, so the `waldur_marketplace_offering` data source reads cleanly.
- 0.2.1's action-body fix (`{"security_groups": [urls]}`) is in upstream.

## Everything else

Resources, data sources and their attributes are upstream's, generated from
Waldur's OpenAPI schema; see `docs/` and the upstream
[changelog](CHANGELOG.md) for the attribute history.

## Licence

MIT, see [LICENSE](LICENSE).
