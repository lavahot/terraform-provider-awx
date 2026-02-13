---
layout: "awx"
page_title: "AWX: awx_project"
sidebar_current: "docs-awx-resource-project"
description: |-
  Manages a project in AWX.
---

# awx_project

Manages a project in AWX. A project represents a collection of Ansible playbooks and related files, typically sourced from a version control system such as Git.

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}

resource "awx_project" "base_service_config" {
  name                 = "base-service-configuration"
  scm_type             = "git"
  scm_url              = "https://github.com/nolte/ansible_playbook-baseline-online-server"
  scm_branch           = "feature/centos8-v2"
  scm_update_on_launch = true
  organization_id      = data.awx_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this project
* `organization_id` - (Required) Numeric ID of the project organization
* `scm_type` - (Required) One of "" (manual), git, hg, svn
* `description` - (Optional) Optional description of this project.
* `local_path` - (Optional) Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.
* `scm_branch` - (Optional) Specific branch, tag or commit to checkout.
* `scm_clean` - (Optional) Whether to remove any local modifications prior to performing an update.
* `scm_credential_id` - (Optional) Numeric ID of the scm used credential
* `scm_delete_on_update` - (Optional) Whether to delete the local project directory before an update.
* `scm_update_cache_timeout` - (Optional) Cache timeout in seconds for SCM update operations.
* `scm_update_on_launch` - (Optional) Whether to update the project from SCM each time a job is launched.
* `scm_url` - (Optional) The URL for the SCM system (e.g., a Git repository URL).

