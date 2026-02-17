---
layout: "awx"
page_title: "AWX: awx_setting"
sidebar_current: "docs-awx-resource-setting"
description: |-
  Manages individual AWX settings.
---

# awx_setting

Manages individual AWX settings. This resource configures generic AWX settings by name and value. Note that deleting this resource only removes it from the Terraform state; it does not reset the setting to its initial value in AWX.

## Example Usage

```hcl
resource "awx_setting" "session_timeout" {
  name  = "SESSION_COOKIE_AGE"
  value = "1800"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the AWX setting to modify.
* `value` - (Required) Value to set for the named setting.
