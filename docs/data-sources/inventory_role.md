---
layout: "awx"
page_title: "AWX: awx_inventory_role"
sidebar_current: "docs-awx-datasource-inventory_role"
description: |-
  Use this data source to look up a role for an inventory in AWX.
---

# awx_inventory_role

Use this data source to look up a role associated with an AWX inventory. This is commonly used to assign inventory-level roles (such as Admin, Use, Update, Adhoc, Read) to teams or users.

## Example Usage

```hcl
resource "awx_inventory" "myinv" {
  name = "My Inventory"
  ...
}

data "awx_inventory_role" "inv_admin_role" {
  name         = "Admin"
  inventory_id = data.awx_inventory.myinv.id
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required) ID of the inventory to reference for inventory roles
* `id` - (Optional) The ID of the role to look up.
* `name` - (Optional) The name of the role to look up (e.g., "Admin", "Use", "Update", "Adhoc", "Read").

