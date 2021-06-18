---
page_title: "metal_spot_market_request Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_spot_market_request`





## Schema

### Required

- **request_id** (String) The id of the Spot Market Request

### Optional

- **id** (String) The ID of this resource.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-only

- **device_ids** (List of String) List of IDs of devices spawned by the referenced Spot Market Request
- **devices_max** (Number) Maximum number devices to be created
- **devices_min** (Number) Miniumum number devices to be created
- **end_at** (String) Date and time When the spot market request will be ended.
- **facilities** (List of String) Facility IDs where devices should be created
- **max_bid_price** (Number) Maximum price user is willing to pay per hour per device
- **metro** (String) Metro where devices should be created.
- **plan** (String) The device plan slug.
- **project_id** (String) Project ID

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)
- **default** (String)
- **delete** (String)
- **update** (String)


