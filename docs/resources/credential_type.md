---
layout: "awx"
page_title: "AWX: awx_credential_type"
sidebar_current: "docs-awx-resource-credential_type"
description: |-
  Manages a credential type in AWX.
---

# awx_credential_type

Manages a custom credential type in AWX. Credential types define the schema for custom credentials, including what input fields they accept and how those fields are injected into playbook execution as extra variables, environment variables, or file-based content.

## Example Usage

```hcl
resource "awx_credential_type" "custom_api" {
  name = "Custom API Token"
  kind = "cloud"
  inputs = jsonencode({
    fields = [{
      id    = "api_token"
      label = "API Token"
      type  = "string"
      secret = true
    }]
  })
  injectors = jsonencode({
    extra_vars = {
      api_token = "{{ api_token }}"
    }
  })
}
```

## Argument Reference

The following arguments are supported:

* `injectors` - (Required) JSON-formatted definition of how the credential fields are injected into job execution (as extra_vars, env, or file content).
* `inputs` - (Required) JSON-formatted definition of the input fields for this credential type.
* `name` - (Required) Name of this credential type.
* `description` - (Optional) Optional description of this credential type.
* `kind` - (Optional) Choices cloud or net

