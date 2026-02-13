---
layout: "awx"
page_title: "AWX: awx_credential"
sidebar_current: "docs-awx-datasource-credential"
description: |-
  Use this data source to query Credential by ID.
---

# awx_credential

Use this data source to query Credential by ID.

## Example Usage

```hcl
data "awx_credential" "machine_cred" {
  id = 1
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the credential to look up.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description of the credential.
* `kind` - The kind of credential (e.g., "cloud", "net", "ssh").
* `name` - The name of the credential.
* `tower_id` - The internal Tower/AWX ID for this credential.
* `username` - The username associated with this credential.
