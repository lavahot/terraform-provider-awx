---
layout: "awx"
page_title: "AWX: awx_job_template_credential"
sidebar_current: "docs-awx-resource-job_template_credential"
description: |-
  Associates a credential with a job template in AWX.
---

# awx_job_template_credential

Associates a credential with a job template in AWX. This resource creates a link between a job template and a credential, allowing the job template to use the credential when running jobs.

## Example Usage

```hcl
resource "awx_job_template_credential" "baseconfig" {
  job_template_id = awx_job_template.baseconfig.id
  credential_id   = awx_credential_machine.pi_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `credential_id` - (Required, ForceNew) The ID of the credential to associate. Changing this forces a new resource to be created.
* `job_template_id` - (Required, ForceNew) The ID of the job template. Changing this forces a new resource to be created.

