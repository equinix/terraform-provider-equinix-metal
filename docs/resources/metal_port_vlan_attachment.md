---
page_title: "metal_port_vlan_attachment Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_port_vlan_attachment`





## Schema

### Required

- **device_id** (String) ID of device to be assigned to the VLAN
- **port_name** (String) Name of network port to be assigned to the VLAN
- **vlan_vnid** (Number) VXLAN Network Identifier, integer

### Optional

- **force_bond** (Boolean) Add port back to the bond when this resource is removed. Default is false
- **id** (String) The ID of this resource.
- **native** (Boolean) Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use depends_on pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above

### Read-only

- **port_id** (String) UUID of device port
- **vlan_id** (String) UUID of VLAN API resource


