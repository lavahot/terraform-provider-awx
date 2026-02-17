---
layout: "awx"
page_title: "AWX: awx_organization"
sidebar_current: "docs-awx-resource-organization"
description: |-
  Manages an organization in AWX.
---

# awx_organization

Manages an organization in AWX. Organizations are logical collections of users, teams, projects, and inventories that provide a means to divide and delegate access.

## Example Usage

```hcl
resource "awx_organization" "default" {
  name            = "acc-test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this organization.
* `custom_virtualenv` - (Optional) Local absolute file path containing a custom Python virtualenv to use
* `description` - (Optional) Optional description of this organization.
* `max_hosts` - (Optional) Maximum number of hosts allowed to be managed by this organization

