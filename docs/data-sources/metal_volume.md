---
page_title: "metal_volume Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_volume`





## Schema

### Optional

- **id** (String) The ID of this resource.
- **name** (String)
- **project_id** (String)
- **volume_id** (String)

### Read-only

- **billing_cycle** (String)
- **created** (String)
- **description** (String)
- **device_ids** (List of String)
- **facility** (String)
- **locked** (Boolean)
- **plan** (String)
- **size** (Number)
- **snapshot_policies** (List of Object) (see [below for nested schema](#nestedatt--snapshot_policies))
- **state** (String)
- **updated** (String)

<a id="nestedatt--snapshot_policies"></a>
### Nested Schema for `snapshot_policies`

Read-only:

- **snapshot_count** (Number)
- **snapshot_frequency** (String)


