---
page_title: "metal_reserved_ip_block Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_reserved_ip_block`





## Schema

### Optional

- **id** (String) ID of the block to look up
- **ip_address** (String) Find block containing this IP address in given project
- **project_id** (String) ID of the project where the searched block should be

### Read-only

- **address** (String)
- **address_family** (Number) 4 or 6
- **cidr** (Number) Length of CIDR prefix of the block as integer
- **cidr_notation** (String) CIDR notation of the looked up block
- **facility** (String) Facility of the block. (for non-global blocks)
- **gateway** (String) IP address of gateway for the block
- **global** (Boolean) Addresses from block are attachable in all locations
- **manageable** (Boolean)
- **management** (Boolean)
- **metro** (String) Metro of the block (for non-global blocks)
- **netmask** (String) Mask in decimal notation, e.g. 255.255.255.0
- **network** (String) Network IP address portion of the block specification
- **public** (Boolean) Addresses from public block are routeable from the Internet
- **quantity** (Number)
- **type** (String) Address type, one of public_ipv4, public_ipv6 and private_ipv4


