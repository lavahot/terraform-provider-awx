---
layout: "awx"
page_title: "AWX: awx_credential_azure_key_vault"
sidebar_current: "docs-awx-resource-credential_azure_key_vault"
description: |-
  Manages an Azure Key Vault credential in AWX.
---

# awx_credential_azure_key_vault

Manages an Azure Key Vault credential in AWX. This credential type is used to configure external credential lookups from Azure Key Vault, allowing AWX to retrieve secrets stored in Azure Key Vault at runtime.

## Example Usage

```hcl
resource "awx_credential_azure_key_vault" "example" {
  name            = "my-azure-kv"
  organization_id = data.awx_organization.default.id
  url             = "https://my-vault.vault.azure.net"
  client          = var.azure_client_id
  secret          = var.azure_client_secret
  tenant          = var.azure_tenant_id
}
```

## Argument Reference

The following arguments are supported:

* `client` - (Required) The Azure Active Directory application client ID.
* `name` - (Required) Name of this credential.
* `organization_id` - (Required) The ID of the organization that this credential belongs to.
* `secret` - (Required) The Azure Active Directory application client secret.
* `tenant` - (Required) The Azure Active Directory tenant ID.
* `url` - (Required) The URL of the Azure Key Vault (e.g., "https://my-vault.vault.azure.net").
* `description` - (Optional) Optional description of this credential.

