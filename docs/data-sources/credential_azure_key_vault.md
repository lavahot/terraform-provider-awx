---
layout: "awx"
page_title: "AWX: awx_credential_azure_key_vault"
sidebar_current: "docs-awx-datasource-credential_azure_key_vault"
description: |-
  Use this data source to look up an Azure Key Vault credential by its ID.
---

# awx_credential_azure_key_vault

Use this data source to look up an Azure Key Vault credential by its ID.

## Example Usage

```hcl
data "awx_credential_azure_key_vault" "example" {
  credential_id = 42
}
```

## Argument Reference

The following arguments are supported:

* `credential_id` - (Required) The ID of the Azure Key Vault credential to look up.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `client` - The Azure Active Directory application client ID.
* `description` - The description of the credential.
* `name` - The name of the credential.
* `organization_id` - The ID of the organization that this credential belongs to.
* `secret` - The Azure Active Directory application client secret.
* `tenant` - The Azure Active Directory tenant ID.
* `url` - The URL of the Azure Key Vault.
