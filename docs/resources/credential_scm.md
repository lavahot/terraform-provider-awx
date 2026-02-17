---
layout: "awx"
page_title: "AWX: awx_credential_scm"
sidebar_current: "docs-awx-resource-credential_scm"
description: |-
  Manages an SCM credential in AWX.
---

# awx_credential_scm

Manages an SCM (Source Control Management) credential in AWX. SCM credentials are used for authenticating with source control systems like Git, Subversion, or Mercurial when syncing project content.

## Example Usage

```hcl
resource "awx_credential_scm" "git_ssh" {
  name            = "my-scm-credential"
  organization_id = data.awx_organization.default.id
  username        = "git"
  ssh_key_data    = file("~/.ssh/id_rsa")
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this credential.
* `organization_id` - (Required) The ID of the organization that this credential belongs to.
* `description` - (Optional) Optional description of this credential.
* `password` - (Optional) The password for SCM authentication (used with HTTPS-based repositories).
* `ssh_key_data` - (Optional) The SSH private key data for SCM authentication.
* `ssh_key_unlock` - (Optional) The passphrase to unlock the SSH private key.
* `username` - (Optional) The username for SCM authentication.

