---
page_title: "metal_ip_block_ranges Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_ip_block_ranges`





## Schema

### Required

- **project_id** (String) ID of the project from which to list the blocks

### Optional

- **facility** (String) Facility code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this and metro, all the block from the project will be listed
- **id** (String) The ID of this resource.
- **metro** (String) Metro code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this and facility, all the block from the project will be listed

### Read-only

- **global_ipv4** (List of String) List of CIDR expressions for Global IPv4 blocks in the project
- **ipv6** (List of String) List of CIDR expressions for IPv6 blocks in the project
- **private_ipv4** (List of String) List of CIDR expressions for Private IPv4 blocks in the project
- **public_ipv4** (List of String) List of CIDR expressions for Public IPv4 blocks in the project


