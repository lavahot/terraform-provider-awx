---
layout: "awx"
page_title: "AWX: awx_credential"
sidebar_current: "docs-awx-resource-credential"
description: |-
  Manages a generic credential in AWX.
---

# awx_credential

Manages a generic credential in AWX. Credentials are used by AWX for authentication when launching jobs against machines, synchronizing with inventory sources, and importing project content from a version control system.

## Example Usage

```hcl
resource "awx_credential" "example" {
  name              = "my-credential"
  credential_type_id = awx_credential_type.custom.id
  organization_id   = data.awx_organization.default.id
  inputs            = jsonencode({
    username = "admin"
    password = "secret"
  })
}
```

## Argument Reference

The following arguments are supported:

* `credential_type_id` - (Required) Specify the type of credential you want to create. Refer to the Ansible Tower documentation for details on each type
* `inputs` - (Required) Credential inputs in JSON format. The expected fields depend on the credential type.
* `name` - (Required) Name of this credential.
* `organization_id` - (Required) The ID of the organization that this credential belongs to.
* `description` - (Optional) Optional description of this credential.

