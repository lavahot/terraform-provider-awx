---
layout: "awx"
page_title: "AWX: awx_organization_role"
sidebar_current: "docs-awx-datasource-organization_role"
description: |-
  Use this data source to look up a role associated with an AWX organization.
---

# awx_organization_role

Use this data source to look up a role associated with an AWX organization. This is commonly used to assign organization-level roles to teams or users.

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}

data "awx_organization_role" "org_admin" {
  name            = "Admin"
  organization_id = data.awx_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required) The ID of the organization to look up roles for
* `id` - (Optional) The ID of the role
* `name` - (Optional) The name of the role (e.g., "Admin", "Read", "Member", "Execute", "Adhoc")

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the role
* `name` - The name of the role
