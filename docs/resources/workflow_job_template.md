---
layout: "awx"
page_title: "AWX: awx_workflow_job_template"
sidebar_current: "docs-awx-resource-workflow_job_template"
description: |-
  Manages a workflow job template in AWX.
---

# awx_workflow_job_template

Manages a workflow job template in AWX. Workflow job templates link together multiple job templates, project syncs, and inventory updates into a single workflow that can be launched as a unit, with conditional logic (success, failure, always) controlling the flow between nodes.

## Example Usage

```hcl
resource "awx_workflow_job_template" "default" {
  name            = "workflow-job"
  organization_id = var.organization_id
  inventory_id    = awx_inventory.default.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this workflow job template. (string, required)
* `allow_simultaneous` - (Optional) Whether multiple workflow jobs from this template can run simultaneously.
* `ask_inventory_on_launch` - (Optional) Whether to prompt for inventory when launching this workflow.
* `ask_limit_on_launch` - (Optional) Whether to prompt for limit when launching this workflow.
* `ask_scm_branch_on_launch` - (Optional) Whether to prompt for SCM branch when launching this workflow.
* `ask_variables_on_launch` - (Optional) Whether to prompt for extra variables when launching this workflow.
* `description` - (Optional) Optional description of this workflow job template.
* `inventory_id` - (Optional) Inventory applied as a prompt, assuming job template prompts for inventory.
* `limit` - (Optional) A host pattern to further constrain the list of managed hosts.
* `organization_id` - (Optional) The organization used to determine access to this template. (id, default=``)
* `scm_branch` - (Optional) The SCM branch to use for this workflow job template.
* `survey_enabled` - (Optional) Whether to enable the survey for this workflow job template.
* `variables` - (Optional) Extra variables for the workflow job template in YAML or JSON format.
* `webhook_credential` - (Optional) The credential to use for webhook authentication.
* `webhook_service` - (Optional) The webhook service (e.g., "github", "gitlab").

