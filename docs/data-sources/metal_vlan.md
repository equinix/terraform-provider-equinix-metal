---
page_title: "metal_vlan Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_vlan`





## Schema

### Optional

- **facility** (String) Facility where the VLAN is deployed
- **id** (String) The ID of this resource.
- **metro** (String) Metro where the VLAN is deployed
- **project_id** (String) ID of parent project of the VLAN. Use together with vxlan and metro or facility
- **vlan_id** (String) Metal UUID of the VLAN resource
- **vxlan** (Number) VXLAN numner of the VLAN. Unique in a project and facility or metro. Use with project_id

### Read-only

- **assigned_devices_ids** (List of String) List of device IDs to which this VLAN is assigned
- **description** (String) VLAN description text


