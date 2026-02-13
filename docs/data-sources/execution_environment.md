---
layout: "awx"
page_title: "AWX: awx_execution_environment"
sidebar_current: "docs-awx-datasource-execution_environment"
description: |-
  Use this data source to look up an execution environment by its ID or name.
---

# awx_execution_environment

Use this data source to look up an execution environment by its ID or name. Execution environments are container images used by AWX to run Ansible automation.

## Example Usage

```hcl
data "awx_execution_environment" "default_ee" {
  name = "Default EE"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the execution environment
* `name` - (Optional) The name of the execution environment

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the execution environment
* `name` - The name of the execution environment
