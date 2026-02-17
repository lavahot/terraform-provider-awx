---
layout: "awx"
page_title: "AWX: awx_job_template_launch"
sidebar_current: "docs-awx-resource-job_template_launch"
description: |-
  Launches an AWX job template.
---

# awx_job_template_launch

Launches an AWX job template. This resource triggers a job run from an existing job template. Note that all arguments use `ForceNew`, meaning any change will destroy and recreate the resource (re-launching the job). This resource is useful for triggering Ansible playbook runs as part of a Terraform workflow.

## Example Usage

```hcl
resource "awx_job_template_launch" "run_baseline" {
  job_template_id     = awx_job_template.baseline.id
  wait_for_completion = true
  extra_vars          = yamlencode({
    target_env = "production"
  })
}
```

## Argument Reference

The following arguments are supported:

* `job_template_id` - (Required, ForceNew) The ID of the job template to launch.
* `limit` - (Optional, ForceNew) A comma-delimited list of hosts to limit the job execution.
* `inventory_id` - (Optional, Computed, ForceNew) Override the inventory for the job launch.
* `extra_vars` - (Optional, ForceNew) Override job template variables. YAML or JSON values are supported.
* `wait_for_completion` - (Optional, Default: `false`, ForceNew) When true, resource creation will wait for the job to complete.
