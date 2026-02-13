---
layout: "awx"
page_title: "AWX: awx_execution_environment"
sidebar_current: "docs-awx-resource-execution_environment"
description: |-
  Manages an Execution Environment in AWX.
---

# awx_execution_environment

Manages an Execution Environment in AWX. Execution environments are container images that include Ansible, collections, Python libraries, and system dependencies needed to run automation. They replace the older concept of custom virtual environments.

## Example Usage

```hcl
resource "awx_execution_environment" "custom_ee" {
  name        = "custom-ee"
  image       = "quay.io/my-org/custom-ee:latest"
  description = "Custom execution environment with additional collections"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this execution environment.
* `image` - (Required) The container image to use for this execution environment (e.g., "quay.io/ansible/awx-ee:latest").
* `description` - (Optional) Description of this execution environment.
* `organization` - (Optional) The organization for this execution environment.
* `credential` - (Optional) The credential to use for pulling the container image (e.g., a container registry credential).
