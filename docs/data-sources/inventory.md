---
layout: "awx"
page_title: "AWX: awx_inventory"
sidebar_current: "docs-awx-datasource-inventory"
description: |-
  Use this data source to look up an inventory in AWX by its ID or name.
---

# awx_inventory

Use this data source to look up an inventory in AWX by its ID or name.

## Example Usage

```hcl
data "awx_inventory" "default" {
  name            = "private_services"
  organization_id = data.awx_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the inventory.
* `name` - (Optional) The name of the inventory.
* `organization_id` - (Optional) The ID of the organization that owns the inventory.

