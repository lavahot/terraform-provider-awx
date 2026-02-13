---
layout: "awx"
page_title: "AWX: awx_credential_machine"
sidebar_current: "docs-awx-resource-credential_machine"
description: |-
  Manages a Machine credential in AWX.
---

# awx_credential_machine

Manages a Machine credential in AWX. Machine credentials are used to authenticate with remote hosts for SSH-based connections when running Ansible playbooks.

## Example Usage

```hcl
resource "awx_credential_machine" "example" {
  name            = "my-machine-credential"
  organization_id = data.awx_organization.default.id
  username        = "ansible"
  ssh_key_data    = file("~/.ssh/id_rsa")
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this credential.
* `organization_id` - (Required) The ID of the organization that this credential belongs to.
* `become_method` - (Optional) The privilege escalation method (e.g., "sudo", "su", "pbrun", "pfexec", "dzdo", "pmrun", "runas").
* `become_password` - (Optional) The password to use for privilege escalation.
* `become_username` - (Optional) The username to use for privilege escalation.
* `description` - (Optional) Optional description of this credential.
* `password` - (Optional) The password for SSH authentication.
* `ssh_key_data` - (Optional) The SSH private key data for authentication.
* `ssh_key_unlock` - (Optional) The passphrase to unlock the SSH private key.
* `ssh_public_key_data` - (Optional) The SSH public key data.
* `username` - (Optional) The username for SSH authentication.

