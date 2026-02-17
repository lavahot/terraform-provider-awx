---
layout: "awx"
page_title: "AWX: awx_schedule"
sidebar_current: "docs-awx-resource-schedule"
description: |-
  Manages a schedule in AWX.
---

# awx_schedule

Manages a schedule in AWX. Schedules define when jobs or workflow jobs should be automatically triggered using iCal recurrence rules (RRULE).

## Example Usage

```hcl
resource "awx_schedule" "default" {
    name                      = "schedule-test"
    rrule                     = "DTSTART;TZID=Europe/Paris:20211214T120000 RRULE:INTERVAL=1;FREQ=DAILY"
    unified_job_template_id   = awx_job_template.baseconfig.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this schedule.
* `unified_job_template_id` - (Required) The ID of the unified job template (job template, workflow, project, or inventory source) to schedule.
* `rrule` - (Required) The iCal recurrence rule (RRULE) that defines the schedule frequency (e.g., "DTSTART;TZID=US/Eastern:20210101T120000 RRULE:INTERVAL=1;FREQ=DAILY").
* `description` - (Optional) Optional description of this schedule.
* `inventory` - (Optional) The ID of the inventory to use when launching the scheduled job (overrides the default).
* `timezone` - (Optional) The timezone for the schedule.
