---
page_title: "metal_ip_attachment Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_ip_attachment`





## Schema

### Required

- **cidr_notation** (String)
- **device_id** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **address** (String)
- **address_family** (Number) Address family as integer (4 or 6)
- **cidr** (Number) Length of CIDR prefix of the block as integer
- **gateway** (String)
- **global** (Boolean) Flag indicating whether IP block is global, i.e. assignable in any location
- **manageable** (Boolean)
- **management** (Boolean)
- **netmask** (String) Mask in decimal notation, e.g. 255.255.255.0
- **network** (String) Network IP address portion of the block specification
- **public** (Boolean) Flag indicating whether IP block is addressable from the Internet


