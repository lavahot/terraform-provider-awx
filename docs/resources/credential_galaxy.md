---
layout: "awx"
page_title: "AWX: awx_credential_galaxy"
sidebar_current: "docs-awx-resource-credential_galaxy"
description: |-
  Manages an Ansible Galaxy / Automation Hub credential in AWX.
---

# awx_credential_galaxy

Manages an Ansible Galaxy / Automation Hub credential in AWX. This credential type is used for authenticating with Ansible Galaxy or a private Automation Hub to download collections and roles.

## Example Usage

```hcl
resource "awx_credential_galaxy" "example" {
  name            = "my-galaxy-credential"
  organization_id = data.awx_organization.default.id
  url             = "https://galaxy.ansible.com/"
  token           = var.galaxy_token
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this credential.
* `organization_id` - (Required) Organization ID for this credential.
* `url` - (Required) The URL of the Galaxy server or Automation Hub.
* `description` - (Optional) Description of this credential.
* `auth_url` - (Optional) The authentication URL for the Galaxy server (e.g., SSO token URL).
* `token` - (Optional, Sensitive) The API token for the Galaxy server or Automation Hub.
