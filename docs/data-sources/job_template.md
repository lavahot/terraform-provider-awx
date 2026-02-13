---
layout: "awx"
page_title: "AWX: awx_job_template"
sidebar_current: "docs-awx-datasource-job_template"
description: |-
  Use this data source to look up a job template in AWX by its ID or name.
---

# awx_job_template

Use this data source to look up a job template in AWX by its ID or name.

## Example Usage

```hcl
data "awx_job_template" "default" {
  name = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the job template.
* `name` - (Optional) The name of the job template.

