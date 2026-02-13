---
layout: "awx"
page_title: "AWX: awx_inventory_group"
sidebar_current: "docs-awx-resource-inventory_group"
description: |-
  Manages an inventory group in AWX.
---

# awx_inventory_group

Manages an inventory group in AWX. Inventory groups are logical collections of hosts within an inventory that allow you to organize and target sets of hosts for automation.

## Example Usage

```hcl
resource "awx_inventory_group" "webservers" {
  name         = "webservers"
  inventory_id = awx_inventory.default.id
  description  = "All web server hosts"
  variables    = yamlencode({
    http_port = 80
  })
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this inventory group.
* `description` - (Optional) Optional description of this inventory group.
* `inventory_id` - (Optional, ForceNew) The ID of the inventory that this group belongs to. Changing this forces a new resource to be created.
* `variables` - (Optional) Group variables in YAML or JSON format.

