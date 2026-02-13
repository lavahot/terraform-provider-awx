---
layout: "awx"
page_title: "AWX: awx_inventory_source"
sidebar_current: "docs-awx-resource-inventory_source"
description: |-
  Manages an inventory source in AWX.
---

# awx_inventory_source

Manages an inventory source in AWX. Inventory sources allow AWX to dynamically fetch inventory data from external sources such as cloud providers, SCM repositories, or other inventory systems.

## Example Usage

```hcl
resource "awx_inventory_source" "scm_source" {
  name              = "scm-inventory-source"
  inventory_id      = awx_inventory.default.id
  source            = "scm"
  source_project_id = awx_project.inventory_project.id
  source_path       = "inventory/hosts.yml"
  update_on_launch  = true
  overwrite         = true
  overwrite_vars    = true
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required, ForceNew) The ID of the inventory that this source belongs to. Changing this forces a new resource to be created.
* `name` - (Required) Name of this inventory source.
* `credential_id` - (Optional) The ID of the credential used to access the inventory source.
* `description` - (Optional) Optional description of this inventory source.
* `enabled_value` - (Optional) Only used if `enabled_var` is set. The host is enabled if the value of the variable specified by `enabled_var` matches this value.
* `enabled_var` - (Optional) If set, AWX will retrieve the enabled state of the host from the given variable.
* `group_by` - (Optional) Limit groups automatically created from inventory source. (Deprecated for some source types.)
* `host_filter` - (Optional) Regular expression to filter hosts returned by the inventory source.
* `instance_filters` - (Optional) Filter expression for cloud inventory sources. (Deprecated for some source types.)
* `overwrite_vars` - (Optional) If true, all host and group variables will be overwritten by the inventory source on each update.
* `overwrite` - (Optional) If true, any hosts and groups that were previously in the inventory but are not found in the updated inventory source will be removed.
* `source_regions` - (Optional) Limit the regions to sync from the cloud inventory source. (Deprecated for some source types.)
* `source_vars` - (Optional) Extra variables for the inventory source in YAML or JSON format.
* `source` - (Optional) The source type for this inventory source (e.g., "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "controller").
* `update_cache_timeout` - (Optional) Time in seconds to consider the inventory source cache valid.
* `update_on_launch` - (Optional) If true, the inventory source will be updated each time a job using this inventory is launched.
* `verbosity` - (Optional) The verbosity level for inventory source updates (0-2).
* `source_project_id` - (Optional) The ID of the project to use as the inventory source (when source is "scm").
* `source_path` - (Optional) The path to the inventory file or script within the source project.

