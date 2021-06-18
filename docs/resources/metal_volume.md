---
page_title: "metal_volume Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_volume`





## Schema

### Required

- **facility** (String)
- **plan** (String)
- **project_id** (String)
- **size** (Number)

### Optional

- **billing_cycle** (String)
- **description** (String)
- **id** (String) The ID of this resource.
- **locked** (Boolean)
- **snapshot_policies** (Block List) (see [below for nested schema](#nestedblock--snapshot_policies))

### Read-only

- **attachments** (List of Object) (see [below for nested schema](#nestedatt--attachments))
- **created** (String)
- **name** (String)
- **state** (String)
- **updated** (String)

<a id="nestedblock--snapshot_policies"></a>
### Nested Schema for `snapshot_policies`

Required:

- **snapshot_count** (Number)
- **snapshot_frequency** (String)


<a id="nestedatt--attachments"></a>
### Nested Schema for `attachments`

Read-only:

- **href** (String)


