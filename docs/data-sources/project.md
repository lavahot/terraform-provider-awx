---
layout: "awx"
page_title: "AWX: awx_project"
sidebar_current: "docs-awx-datasource-project"
description: |-
  Use this data source to look up a project in AWX by its ID or name.
---

# awx_project

Use this data source to look up a project in AWX by its ID or name.

## Example Usage

```hcl
data "awx_project" "default" {
  name = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the project.
* `name` - (Optional) The name of the project.

