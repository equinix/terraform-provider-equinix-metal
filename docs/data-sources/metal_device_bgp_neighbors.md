---
page_title: "metal_device_bgp_neighbors Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_device_bgp_neighbors`





## Schema

### Required

- **device_id** (String) UUID of BGP-enabled device whose neighbors to list

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **bgp_neighbors** (List of Object) Array of BGP neighbor records (see [below for nested schema](#nestedatt--bgp_neighbors))

<a id="nestedatt--bgp_neighbors"></a>
### Nested Schema for `bgp_neighbors`

Read-only:

- **address_family** (Number)
- **customer_as** (Number)
- **customer_ip** (String)
- **md5_enabled** (Boolean)
- **md5_password** (String)
- **multihop** (Boolean)
- **peer_as** (Number)
- **peer_ips** (List of String)
- **routes_in** (List of Object) (see [below for nested schema](#nestedobjatt--bgp_neighbors--routes_in))
- **routes_out** (List of Object) (see [below for nested schema](#nestedobjatt--bgp_neighbors--routes_out))

<a id="nestedobjatt--bgp_neighbors--routes_in"></a>
### Nested Schema for `bgp_neighbors.routes_in`

Read-only:

- **exact** (Boolean)
- **route** (String)


<a id="nestedobjatt--bgp_neighbors--routes_out"></a>
### Nested Schema for `bgp_neighbors.routes_out`

Read-only:

- **exact** (Boolean)
- **route** (String)


