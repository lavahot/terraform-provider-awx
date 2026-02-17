---
layout: "awx"
page_title: "AWX: awx_inventory_group"
sidebar_current: "docs-awx-datasource-inventory_group"
description: |-
  Use this data source to look up an inventory group in AWX.
---

# awx_inventory_group

Use this data source to look up an inventory group in AWX by its name or ID within a specific inventory.

## Example Usage

```hcl
data "awx_inventory_group" "default" {
  name         = "k3sPrimary"
  inventory_id = data.awx_inventory.default.id
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required) The ID of the inventory that this group belongs to.
* `id` - (Optional) The ID of the inventory group.
* `name` - (Optional) The name of the inventory group.

