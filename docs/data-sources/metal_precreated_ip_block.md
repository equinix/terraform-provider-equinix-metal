---
page_title: "metal_precreated_ip_block Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_precreated_ip_block`





## Schema

### Required

- **address_family** (Number) 4 or 6, depending on which block you are looking for.
- **project_id** (String) ID of the project where the searched block should be.
- **public** (Boolean) Whether to look for public or private block.

### Optional

- **facility** (String) Facility of the searched block. (for non-global blocks).
- **global** (Boolean) Whether to look for global block. Default is false for backward compatibility.
- **id** (String) The ID of this resource.
- **metro** (String) Metro of the searched block (for non-global blocks).

### Read-only

- **address** (String)
- **cidr** (Number) Length of CIDR prefix of the block as integer
- **cidr_notation** (String) CIDR notation of the looked up block.
- **gateway** (String)
- **manageable** (Boolean)
- **management** (Boolean)
- **netmask** (String) Mask in decimal notation, e.g. 255.255.255.0
- **network** (String) Network IP address portion of the block specification
- **quantity** (Number)
- **type** (String)


