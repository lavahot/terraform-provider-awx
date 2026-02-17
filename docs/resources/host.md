---
layout: "awx"
page_title: "AWX: awx_host"
sidebar_current: "docs-awx-resource-host"
description: |-
  Manages a host in AWX.
---

# awx_host

Manages a host in an AWX inventory. Hosts represent the systems that AWX can manage and run automation against.

## Example Usage

```hcl
resource "awx_host" "k3snode1" {
  name         = "k3snode1"
  description  = "pi node 1"
  inventory_id = data.awx_inventory.default.id
  group_ids = [
    data.awx_inventory_group.default.id,
    data.awx_inventory_group.pinodes.id,
  ]
  enabled   = true
  variables = <<YAML
---
ansible_host: 192.168.178.29
YAML
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required) The ID of the inventory that this host belongs to.
* `name` - (Required) The name or hostname of this host.
* `description` - (Optional) Optional description of this host.
* `enabled` - (Optional) Whether this host is enabled and can be used for job runs.
* `group_ids` - (Optional) A list of inventory group IDs that this host belongs to.
* `instance_id` - (Optional) The instance ID for this host (used for cloud-based inventories).
* `variables` - (Optional) Host variables in YAML or JSON format.

