---
page_title: "metal_spot_market_request Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_spot_market_request`





## Schema

### Required

- **devices_max** (Number) Maximum number devices to be created
- **devices_min** (Number) Miniumum number devices to be created
- **instance_parameters** (Block List, Min: 1, Max: 1) Parameters for devices provisioned from this request. You can find the parameter description from the [metal_device doc](device.md) (see [below for nested schema](#nestedblock--instance_parameters))
- **max_bid_price** (Number) Maximum price user is willing to pay per hour per device
- **project_id** (String) Project ID

### Optional

- **facilities** (List of String) Facility IDs where devices should be created
- **id** (String) The ID of this resource.
- **metro** (String) Metro where devices should be created
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **wait_for_devices** (Boolean) On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed

<a id="nestedblock--instance_parameters"></a>
### Nested Schema for `instance_parameters`

Required:

- **billing_cycle** (String)
- **hostname** (String)
- **operating_system** (String)
- **plan** (String)

Optional:

- **always_pxe** (Boolean)
- **customdata** (String)
- **description** (String)
- **features** (List of String)
- **ipxe_script_url** (String)
- **locked** (Boolean)
- **project_ssh_keys** (List of String)
- **tags** (List of String)
- **user_ssh_keys** (List of String)
- **userdata** (String)

Read-only:

- **termintation_time** (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)
- **delete** (String)


