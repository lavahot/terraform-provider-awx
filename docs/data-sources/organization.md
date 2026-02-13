---
layout: "awx"
page_title: "AWX: awx_organization"
sidebar_current: "docs-awx-datasource-organization"
description: |-
  Use this data source to look up an organization in AWX by its ID or name.
---

# awx_organization

Use this data source to look up an organization in AWX by its ID or name.

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the organization.
* `name` - (Optional) The name of the organization.

