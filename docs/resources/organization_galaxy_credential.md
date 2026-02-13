---
layout: "awx"
page_title: "AWX: awx_organization_galaxy_credential"
sidebar_current: "docs-awx-resource-organization_galaxy_credential"
description: |-
  Associates a Galaxy credential with an AWX organization.
---

# awx_organization_galaxy_credential

Associates a Galaxy credential with an AWX organization. This allows the organization to use the specified Galaxy or Automation Hub credential for downloading Ansible collections and roles.

## Example Usage

```hcl
resource "awx_organization_galaxy_credential" "example" {
  organization_id = awx_organization.example.id
  credential_id   = awx_credential_galaxy.hub.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required, ForceNew) The ID of the organization.
* `credential_id` - (Required, ForceNew) The ID of the Galaxy credential to associate.
