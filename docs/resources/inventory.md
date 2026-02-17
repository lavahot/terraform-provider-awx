---
layout: "awx"
page_title: "AWX: awx_inventory"
sidebar_current: "docs-awx-resource-inventory"
description: |-
  Manages an inventory in AWX.
---

# awx_inventory

Manages an inventory in AWX. An inventory is a collection of hosts and groups that Ansible automation can be run against.

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}

resource "awx_inventory" "default" {
  name            = "acc-test"
  organization_id = data.awx_organization.default.id
  variables       = <<YAML
---
system_supporters:
  - pi
YAML
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this inventory.
* `organization_id` - (Required) The ID of the organization that this inventory belongs to.
* `description` - (Optional) Optional description of this inventory.
* `host_filter` - (Optional) Filter expression to limit hosts included in the inventory (for smart inventories).
* `kind` - (Optional) The kind of inventory. Set to empty string for a normal inventory or "smart" for a smart inventory.
* `variables` - (Optional) Inventory variables in YAML or JSON format.

