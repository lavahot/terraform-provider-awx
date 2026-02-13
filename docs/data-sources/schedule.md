---
layout: "awx"
page_title: "AWX: awx_schedule"
sidebar_current: "docs-awx-datasource-schedule"
description: |-
  Use this data source to look up a schedule in AWX by its ID or name.
---

# awx_schedule

Use this data source to look up a schedule in AWX by its ID or name.

## Example Usage

```hcl
data "awx_schedule" "default" {
  name            = "private_services"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the schedule.
* `name` - (Optional) The name of the schedule.

