---
layout: "awx"
page_title: "AWX: awx_organizations"
sidebar_current: "docs-awx-datasource-organizations"
description: |-
  Use this data source to list all organizations in AWX.
---

# awx_organizations

Use this data source to list all organizations in AWX. Returns a list of all organizations with their IDs and names.

## Example Usage

```hcl
data "awx_organizations" "all" {}

output "org_names" {
  value = data.awx_organizations.all.organizations[*].name
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `organizations` - A list of all organizations. Each element has:
  * `id` - The ID of the organization
  * `name` - The name of the organization
