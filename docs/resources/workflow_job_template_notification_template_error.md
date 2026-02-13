---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_notification_template_error"
sidebar_current: "docs-awx-resource-workflow_job_template_notification_template_error"
description: |-
  Associates a notification template with a workflow job template to send notifications on error.
---

# awx_workflow_job_template_notification_template_error

Associates a notification template with a workflow job template to send notifications on error. When the workflow job fails, the associated notification template will be triggered.

## Example Usage

```hcl
resource "awx_workflow_job_template_notification_template_error" "example" {
  workflow_job_template_id = awx_workflow_job_template.deploy.id
  notification_template_id = awx_notification_template.slack_alerts.id
}
```

## Argument Reference

The following arguments are supported:

* `workflow_job_template_id` - (Required, ForceNew) The ID of the workflow job template.
* `notification_template_id` - (Required, ForceNew) The ID of the notification template to associate.
