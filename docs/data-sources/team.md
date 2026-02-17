---
layout: "awx"
page_title: "AWX: awx_team"
sidebar_current: "docs-awx-datasource-team"
description: |-
  Use this data source to look up a team in AWX by its ID or name.
---

# awx_team

Use this data source to look up a team in AWX by its ID or name.

## Example Usage

```hcl
data "awx_team" "default" {
  name = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the team.
* `name` - (Optional) The name of the team.

