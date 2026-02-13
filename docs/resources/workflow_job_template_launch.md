---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_launch"
sidebar_current: "docs-awx-resource-workflow_job_template_launch"
description: |-
  Launches an AWX workflow job template.
---

# awx_workflow_job_template_launch

Launches an AWX workflow job template. This resource triggers a workflow run from an existing workflow job template. Note that all arguments use `ForceNew`, meaning any change will destroy and recreate the resource (re-launching the workflow). This is useful for triggering multi-step automation workflows as part of a Terraform workflow.

## Example Usage

```hcl
resource "awx_workflow_job_template_launch" "deploy" {
  workflow_job_template_id = awx_workflow_job_template.deploy_workflow.id
  wait_for_completion      = true
  extra_vars               = yamlencode({
    release_version = "1.2.3"
  })
}
```

## Argument Reference

The following arguments are supported:

* `workflow_job_template_id` - (Required, ForceNew) The ID of the workflow job template to launch.
* `extra_vars` - (Optional, ForceNew) Override workflow job template variables. YAML or JSON values are supported.
* `wait_for_completion` - (Optional, Default: `false`, ForceNew) When true, resource creation will wait for the workflow job to complete.
