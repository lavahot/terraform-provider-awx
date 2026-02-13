---
layout: "awx"
page_title: "AWX: awx_instance_group"
sidebar_current: "docs-awx-resource-instance_group"
description: |-
  Manages an Instance Group in AWX.
---

# awx_instance_group

Manages an Instance Group in AWX. Instance groups are used to group automation controller instances and direct work to specific sets of instances. Container groups are a special type of instance group that run jobs in Kubernetes pods.

## Example Usage

```hcl
resource "awx_instance_group" "container_group" {
  name               = "my-container-group"
  is_container_group = true
  pod_spec_override  = yamlencode({
    apiVersion = "v1"
    kind       = "Pod"
    metadata = {
      namespace = "awx"
    }
    spec = {
      containers = [{
        image = "quay.io/ansible/awx-ee:latest"
        name  = "worker"
      }]
    }
  })
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this instance group.
* `is_container_group` - (Optional, Default: `true`) Whether this is a container group.
* `policy_instance_minimum` - (Optional, Default: `0`) Minimum number of instances to keep in this group.
* `policy_instance_percentage` - (Optional, Default: `0`) Percentage of instances to allocate to this group.
* `pod_spec_override` - (Optional) Custom pod specification override in YAML or JSON format, used when `is_container_group` is true.
