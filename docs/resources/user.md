---
layout: "awx"
page_title: "AWX: awx_user"
sidebar_current: "docs-awx-resource-user"
description: |-
  Manages a user in AWX.
---

# awx_user

Manages a user in AWX. This resource allows you to create and manage AWX user accounts, including setting role-based access control entitlements.

## Example Usage

```hcl
resource "awx_user" "developer" {
  username          = "jdoe"
  password          = var.user_password
  first_name        = "John"
  last_name         = "Doe"
  email             = "jdoe@example.com"
  is_superuser      = false
  is_system_auditor = false
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) The username for this user.
* `password` - (Required, Sensitive) The password for this user.
* `first_name` - (Optional) The first name of this user.
* `last_name` - (Optional) The last name of this user.
* `email` - (Optional) The email address of this user.
* `is_superuser` - (Optional) Whether this user is a superuser (administrator).
* `is_system_auditor` - (Optional) Whether this user is a system auditor.
* `role_entitlement` - (Optional) Set of role IDs of the role entitlements for this user.
