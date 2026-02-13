---
layout: "awx"
page_title: "AWX: awx_workflow_job_template"
sidebar_current: "docs-awx-datasource-workflow_job_template"
description: |-
  Use this data source to look up a workflow job template in AWX by its ID or name.
---

# awx_workflow_job_template

Use this data source to look up a workflow job template in AWX by its ID or name.

## Example Usage

```hcl
data "awx_workflow_job_template" "default" {
  name = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the workflow job template.
* `name` - (Optional) The name of the workflow job template.

