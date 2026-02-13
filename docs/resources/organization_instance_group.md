---
layout: "awx"
page_title: "AWX: awx_organization_instance_group"
sidebar_current: "docs-awx-resource-organization_instance_group"
description: |-
  Associates an instance group with an AWX organization.
---

# awx_organization_instance_group

Associates an instance group with an AWX organization. This resource creates a link between an organization and an instance group so that jobs run by the organization will be routed to instances in the associated instance group.

## Example Usage

```hcl
resource "awx_organization_instance_group" "example" {
  organization_id   = awx_organization.example.id
  instance_group_id = awx_instance_group.container_group.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required, ForceNew) The ID of the organization.
* `instance_group_id` - (Required, ForceNew) The ID of the instance group to associate.
