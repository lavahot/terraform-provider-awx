---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_node"
sidebar_current: "docs-awx-resource-workflow_job_template_node"
description: |-
  Manages a workflow job template node in AWX.
---

# awx_workflow_job_template_node

Manages a workflow job template node in AWX. Workflow nodes represent individual steps in a workflow job template. Each node is linked to a unified job template (such as a job template, project sync, or inventory update) and can have success, failure, and always paths to other nodes.

## Example Usage

```hcl
resource "random_uuid" "workflow_node_base_uuid" {}

resource "awx_workflow_job_template_node" "default" {
  workflow_job_template_id = awx_workflow_job_template.default.id
  unified_job_template_id  = awx_job_template.baseconfig.id
  inventory_id             = awx_inventory.default.id
  identifier               = random_uuid.workflow_node_base_uuid.result
}
```

## Argument Reference

The following arguments are supported:

* `identifier` - (Required) A unique identifier for this node within the workflow.
* `unified_job_template_id` - (Required) The ID of the unified job template (job template, project, or inventory source) to run for this node.
* `workflow_job_template_id` - (Required) The ID of the workflow job template that this node belongs to.
* `all_parents_must_converge` - (Optional) Whether all parent nodes must complete successfully before this node runs.
* `diff_mode` - (Optional) Whether to enable diff mode for this node.
* `extra_data` - (Optional) Extra data for this node in YAML or JSON format.
* `inventory_id` - (Optional) Inventory applied as a prompt, assuming job template prompts for inventory.
* `job_tags` - (Optional) Tags to apply when running this node.
* `job_type` - (Optional) The job type for this node (e.g., "run", "check").
* `limit` - (Optional) A host pattern to limit which hosts are affected by this node.
* `scm_branch` - (Optional) The SCM branch to use for this node.
* `skip_tags` - (Optional) Tags to skip when running this node.
* `verbosity` - (Optional) The verbosity level for this node (0-5).

