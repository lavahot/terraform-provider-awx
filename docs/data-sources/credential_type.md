---
layout: "awx"
page_title: "AWX: awx_credential_type"
sidebar_current: "docs-awx-datasource-credential_type"
description: |-
  Use this data source to query Credential Type by ID.
---

# awx_credential_type

Use this data source to query Credential Type by ID.

## Example Usage

```hcl
data "awx_credential_type" "machine" {
  id = 1
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the credential type to look up.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description of the credential type.
* `injectors` - JSON-formatted definition of how credential fields are injected.
* `inputs` - JSON-formatted definition of input fields for this credential type.
* `kind` - The kind of credential type ("cloud" or "net").
* `name` - The name of the credential type.
