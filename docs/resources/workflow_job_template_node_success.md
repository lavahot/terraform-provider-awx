---
layout: "awx"
page_title: "AWX: awx_workflow_job_template_node_success"
sidebar_current: "docs-awx-resource-workflow_job_template_node_success"
description: |-
  Manages a 'success' path node in an AWX workflow job template.
---

# awx_workflow_job_template_node_success

Manages a 'success' path node in an AWX workflow job template. A success-path node will only execute when the parent node completes successfully. This is used to chain dependent automation steps.

## Example Usage

```hcl
resource "random_uuid" "workflow_node_k3s_uuid" {}

resource "awx_workflow_job_template_node_success" "k3s" {
  workflow_job_template_node_id = awx_workflow_job_template_node.default.id
  unified_job_template_id       = awx_job_template.k3s.id
  inventory_id                  = awx_inventory.default.id
  identifier                    = random_uuid.workflow_node_k3s_uuid.result
}
```

## Argument Reference

The following arguments are supported:

* `identifier` - (Required) A unique identifier for this node within the workflow.
* `unified_job_template_id` - (Required) The ID of the unified job template to run for this node.
* `workflow_job_template_node_id` - (Required) The ID of the parent workflow node that this success-path node follows.
* `all_parents_must_converge` - (Optional) Whether all parent nodes must complete before this node runs.
* `diff_mode` - (Optional) Whether to enable diff mode for this node.
* `extra_data` - (Optional) Extra data for this node in YAML or JSON format.
* `inventory_id` - (Optional) Inventory applied as a prompt, assuming job template prompts for inventory.
* `job_tags` - (Optional) Tags to apply when running this node.
* `job_type` - (Optional) The job type for this node (e.g., "run", "check").
* `limit` - (Optional) A host pattern to limit which hosts are affected by this node.
* `scm_branch` - (Optional) The SCM branch to use for this node.
* `skip_tags` - (Optional) Tags to skip when running this node.
* `verbosity` - (Optional) The verbosity level for this node (0-5).

