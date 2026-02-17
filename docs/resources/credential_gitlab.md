---
layout: "awx"
page_title: "AWX: awx_credential_gitlab"
sidebar_current: "docs-awx-resource-credential_gitlab"
description: |-
  Manages a GitLab personal access token credential in AWX.
---

# awx_credential_gitlab

Manages a GitLab personal access token credential in AWX. This credential type is used for authenticating with GitLab APIs, such as when syncing projects from GitLab SCM.

## Example Usage

```hcl
resource "awx_credential_gitlab" "example" {
  name            = "my-gitlab-token"
  organization_id = data.awx_organization.default.id
  token           = var.gitlab_token
  description     = "GitLab PAT for project sync"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this credential.
* `token` - (Required, Sensitive) The GitLab personal access token.
* `organization_id` - (Optional) Organization ID for this credential.
* `description` - (Optional) Description of this credential.
