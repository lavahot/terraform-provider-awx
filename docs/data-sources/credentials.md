---
layout: "awx"
page_title: "AWX: awx_credentials"
sidebar_current: "docs-awx-datasource-credentials"
description: |-
  Use this data source to list all credentials in AWX.
---

# awx_credentials

Use this data source to list all credentials in AWX.

## Example Usage

```hcl
data "awx_credentials" "all" {}

output "credential_names" {
  value = data.awx_credentials.all.credentials[*].name
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `credentials` - A list of all credentials.
  * `description` - The description of the credential.
  * `id` - The ID of the credential.
  * `kind` - The kind of credential.
  * `name` - The name of the credential.
  * `username` - The username associated with the credential.
