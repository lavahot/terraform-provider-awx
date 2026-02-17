---
layout: "awx"
page_title: "AWX: awx_job_template_notification_template_started"
sidebar_current: "docs-awx-resource-job_template_notification_template_started"
description: |-
  Associates a notification template with a job template to send notifications when a job starts.
---

# awx_job_template_notification_template_started

Associates a notification template with a job template to send notifications when a job starts. When a job launched from this template begins execution, the associated notification template will be triggered.

## Example Usage

```hcl
resource "awx_job_template_notification_template_started" "baseconfig" {
    job_template_id            = awx_job_template.baseconfig.id
    notification_template_id   = awx_notification_template.default.id
}
```

## Argument Reference

The following arguments are supported:

* `notification_template_id` - (Required, ForceNew) The ID of the notification template. Changing this forces a new resource to be created.
* `job_template_id` - (Required, ForceNew) The ID of the job template. Changing this forces a new resource to be created.

