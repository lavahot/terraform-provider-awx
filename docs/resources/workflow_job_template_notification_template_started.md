---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_notification_template_started"
sidebar_current: "docs-awx-resource-workflow_job_template_notification_template_started"
description: |-
  Associates a notification template with a workflow job template to send notifications when the workflow starts.
---

# awx_workflow_job_template_notification_template_started

Associates a notification template with a workflow job template to send notifications when the workflow starts. When the workflow job begins execution, the associated notification template will be triggered.

## Example Usage

```hcl
resource "awx_workflow_job_template_notification_template_started" "example" {
  workflow_job_template_id = awx_workflow_job_template.deploy.id
  notification_template_id = awx_notification_template.slack_alerts.id
}
```

## Argument Reference

The following arguments are supported:

* `workflow_job_template_id` - (Required, ForceNew) The ID of the workflow job template.
* `notification_template_id` - (Required, ForceNew) The ID of the notification template to associate.
