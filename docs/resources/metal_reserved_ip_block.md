---
page_title: "metal_reserved_ip_block Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_reserved_ip_block`





## Schema

### Required

- **project_id** (String) The metal project ID where to allocate the address block
- **quantity** (Number) The number of allocated /32 addresses, a power of 2

### Optional

- **description** (String) Arbitrary description
- **facility** (String) Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with metro
- **id** (String) The ID of this resource.
- **metro** (String) Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with facility
- **type** (String) Either global_ipv4 or public_ipv4, defaults to public_ipv4 for backward compatibility

### Read-only

- **address** (String)
- **address_family** (Number) Address family as integer (4 or 6)
- **cidr** (Number) Length of CIDR prefix of the block as integer
- **cidr_notation** (String)
- **gateway** (String)
- **global** (Boolean) Flag indicating whether IP block is global, i.e. assignable in any location
- **manageable** (Boolean)
- **management** (Boolean)
- **netmask** (String) Mask in decimal notation, e.g. 255.255.255.0
- **network** (String) Network IP address portion of the block specification
- **public** (Boolean) Flag indicating whether IP block is addressable from the Internet


