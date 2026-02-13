---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_schedule"
sidebar_current: "docs-awx-resource-workflow_job_template_schedule"
description: |-
  Manages a schedule for a workflow job template in AWX.
---

# awx_workflow_job_template_schedule

Manages a schedule for a workflow job template in AWX. This resource allows you to define recurring schedules for workflow job templates using iCal recurrence rules (RRULE).

## Example Usage

```hcl
resource "awx_workflow_job_template_schedule" "default" {
    workflow_job_template_id  = awx_workflow_job_template.default.id

    name                      = "schedule-test"
    rrule                     = "DTSTART;TZID=Europe/Paris:20211214T120000 RRULE:INTERVAL=1;FREQ=DAILY"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this schedule.
* `rrule` - (Required) The iCal recurrence rule (RRULE) that defines the schedule frequency.
* `workflow_job_template_id` - (Required) The ID of the workflow job template to schedule.
* `description` - (Optional) Optional description of this schedule.
* `inventory` - (Optional) The ID of the inventory to use when launching the scheduled workflow (overrides the default).
* `timezone` - (Optional) The timezone for the schedule.
* `extra_data` - (Optional) Extra data for the scheduled workflow in YAML format.
