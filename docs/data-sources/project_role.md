---
layout: "awx"
page_title: "AWX: awx_project_role"
sidebar_current: "docs-awx-datasource-project_role"
description: |-
  Use this data source to look up a role for a project in AWX.
---

# awx_project_role

Use this data source to look up a role associated with an AWX project. This is commonly used to assign project-level roles (such as Admin, Use, Update, Read) to teams or users.

## Example Usage

```hcl
resource "awx_project" "myproj" {
  name = "My AWX Project"
  ...
}

data "awx_project_role" "proj_admins" {
  name       = "Admin"
  project_id = resource.awx_project.myproj.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project to reference for the named role
* `id` - (Optional)
* `name` - (Optional)

