---
layout: "awx"
page_title: "AWX: awx_job_template"
sidebar_current: "docs-awx-resource-job_template"
description: |-
  Manages a job template in AWX.
---

# awx_job_template

Manages a job template in AWX. A job template is a definition and set of parameters for running an Ansible playbook. Job templates encourage reuse of Ansible playbook content and allow job runs to be delegated to users who need to execute them.

## Example Usage

```hcl
data "awx_inventory" "default" {
  name            = "private_services"
  organization_id = data.awx_organization.default.id
}

resource "awx_job_template" "baseconfig" {
  name           = "baseconfig"
  job_type       = "run"
  inventory_id   = data.awx_inventory.default.id
  project_id     = awx_project.base_service_config.id
  playbook       = "master-configure-system.yml"
  become_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required) The ID of the inventory to use for this job template.
* `job_type` - (Required) One of: run, check, scan
* `name` - (Required) Name of this job template.
* `project_id` - (Required) The ID of the project containing the playbook for this job template.
* `allow_simultaneous` - (Optional) Whether multiple jobs from this template can run simultaneously.
* `ask_credential_on_launch` - (Optional) Whether to prompt for credentials when launching jobs from this template.
* `ask_diff_mode_on_launch` - (Optional) Whether to prompt for diff mode when launching jobs from this template.
* `ask_inventory_on_launch` - (Optional) Whether to prompt for inventory when launching jobs from this template.
* `ask_job_type_on_launch` - (Optional) Whether to prompt for job type when launching jobs from this template.
* `ask_limit_on_launch` - (Optional) Whether to prompt for limit when launching jobs from this template.
* `ask_skip_tags_on_launch` - (Optional) Whether to prompt for skip tags when launching jobs from this template.
* `ask_tags_on_launch` - (Optional) Whether to prompt for job tags when launching jobs from this template.
* `ask_variables_on_launch` - (Optional) Whether to prompt for extra variables when launching jobs from this template.
* `ask_verbosity_on_launch` - (Optional) Whether to prompt for verbosity when launching jobs from this template.
* `become_enabled` - (Optional) Whether to enable privilege escalation (become) for this job template.
* `custom_virtualenv` - (Optional) Path to a custom Python virtual environment to use.
* `description` - (Optional) Optional description of this job template.
* `diff_mode` - (Optional) Whether to enable diff mode for this job template, showing changes made to files.
* `extra_vars` - (Optional) Extra variables for the job template in YAML or JSON format.
* `force_handlers` - (Optional) Whether to force notified handler execution even if a task fails.
* `forks` - (Optional) The number of parallel processes to use for playbook execution.
* `host_config_key` - (Optional) The host config key for callback-based provisioning.
* `job_tags` - (Optional) Comma-separated list of tags to apply, limiting the playbook run to matching tagged tasks.
* `limit` - (Optional) A host pattern to further constrain the list of managed hosts.
* `playbook` - (Optional) The playbook file to execute from the project.
* `skip_tags` - (Optional) Comma-separated list of tags to skip when running the playbook.
* `start_at_task` - (Optional) The task name to start the playbook execution at.
* `survey_enabled` - (Optional) Whether to enable the survey for this job template.
* `timeout` - (Optional) Maximum time in seconds to allow the job to run (0 means no timeout).
* `use_fact_cache` - (Optional) Whether to enable fact caching for this job template.
* `verbosity` - (Optional) One of 0,1,2,3,4,5

