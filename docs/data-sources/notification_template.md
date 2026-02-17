---
layout: "awx"
page_title: "AWX: awx_notification_template"
sidebar_current: "docs-awx-datasource-notification-template"
description: |-
  Use this data source to look up a notification template in AWX by its ID or name.
---

# awx_notification_template

Use this data source to look up a notification template in AWX by its ID or name.

## Example Usage

```hcl
data "awx_notification_template" "default" {
  name            = "private_services"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the notification template.
* `name` - (Optional) The name of the notification template.

