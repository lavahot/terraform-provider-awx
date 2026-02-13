---
layout: "awx"
page_title: "AWX: awx_credential_google_compute_engine"
sidebar_current: "docs-awx-resource-credential_google_compute_engine"
description: |-
  Manages a Google Compute Engine credential in AWX.
---

# awx_credential_google_compute_engine

Manages a Google Compute Engine credential in AWX. This credential type is used for authenticating with Google Cloud Platform for dynamic inventory and other GCP-related operations.

## Example Usage

```hcl
resource "awx_credential_google_compute_engine" "example" {
  name            = "my-gce-credential"
  organization_id = data.awx_organization.default.id
  username        = "my-service-account@my-project.iam.gserviceaccount.com"
  project         = "my-gcp-project"
  ssh_key_data    = file("~/.ssh/gce_key")
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this credential.
* `organization_id` - (Required) The ID of the organization that this credential belongs to.
* `project` - (Required) The GCP project ID associated with this credential.
* `ssh_key_data` - (Required) The SSH private key or service account JSON key data.
* `username` - (Required) The service account email address for GCE.
* `description` - (Optional) Optional description of this credential.

