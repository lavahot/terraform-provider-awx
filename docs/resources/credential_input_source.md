---
layout: "awx"
page_title: "AWX: awx_credential_input_source"
sidebar_current: "docs-awx-resource-credential_input_source"
description: |-
  Manages a credential input source in AWX.
---

# awx_credential_input_source

Manages a credential input source in AWX. Credential input sources allow credentials to dynamically pull their values from an external secret management system (such as Azure Key Vault, HashiCorp Vault, or CyberArk) at runtime.

## Example Usage

```hcl
resource "awx_credential_input_source" "example" {
  input_field_name = "password"
  target           = awx_credential_machine.example.id
  source           = awx_credential_azure_key_vault.example.id
  metadata = {
    secret_field = "my-secret-name"
  }
}
```

## Argument Reference

The following arguments are supported:

* `input_field_name` - (Required) The name of the input field on the target credential that will be populated from the external source.
* `source` - (Required) The ID of the source credential that provides the external lookup (e.g., Azure Key Vault credential).
* `target` - (Required) The ID of the target credential whose input field will be dynamically populated.
* `description` - (Optional) Optional description of this credential input source.
* `metadata` - (Optional) A map of metadata key-value pairs specific to the source credential type (e.g., secret name, secret version).

