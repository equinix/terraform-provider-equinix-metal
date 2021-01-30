---
page_title: "metal_volume Resource - terraform-provider-metal"
subcategory: ""
description: |-
  Provides an Equinix Metal Block Storage Volume resource to allow you to manage block volumes on your account.  Once created by Terraform, they must then be attached and mounted using the api and metal_block_attach and metal_block_detach scripts.
---

# Resource `metal_volume`

Provides an Equinix Metal Block Storage Volume resource to allow you to manage block volumes on your account.  Once created by Terraform, they must then be attached and mounted using the api and `metal_block_attach` and `metal_block_detach` scripts.

## Example Usage

```terraform
locals {
    project_id = "552345b2-ee46-4673-93a8-de2c2bdba33b"
}


resource "metal_volume" "volume1" {
  description   = "terraform-volume-1"
  facility      = "ewr1"
  project_id    = local.project_id
  plan          = "storage_1"
  size          = 100
  billing_cycle = "hourly"

  snapshot_policies {
    snapshot_frequency = "1day"
    snapshot_count     = 7
  }

  snapshot_policies {
    snapshot_frequency = "1month"
    snapshot_count     = 6
  }
}
```

## Schema

### Required

- **facility** (String) The facility in which the resource should be created
- **plan** (String) Plan/Performance tier
- **project_id** (String) ID of metap project in which the resource should be created
- **size** (Number) Volume size in GB. Must be at elast 100

### Optional

- **billing_cycle** (String) hourly (default) or monthly
- **description** (String) Optional description for the volume
- **id** (String) The ID of this resource.
- **locked** (Boolean) Setting this parameter will prevent resource deletion
- **snapshot_policies** (Block List) List of policies for volume snapshots (see [below for nested schema](#nestedblock--snapshot_policies))

### Read-only

- **attachments** (List of Object) List of IDs of device-volume attachment API resources linked to this volume (see [below for nested schema](#nestedatt--attachments))
- **created** (String) Creation timestamp
- **name** (String) The generated name of the volume
- **state** (String) Provisioning state of the resource
- **updated** (String) Update timestamp

<a id="nestedblock--snapshot_policies"></a>
### Nested Schema for `snapshot_policies`

Required:

- **snapshot_count** (Number) Number of snapshots to keep
- **snapshot_frequency** (String) How often to make snapshots, e.g "1day"


<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Read-only:

- **href** (String)

## Import

Import is supported using the following syntax:

```shell
# import existing volume
terraform import metal_volume.test 07af1eec-091e-4ddd-a4ed-9b85928d4c36
```
