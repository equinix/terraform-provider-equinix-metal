---
page_title: "metal_virtual_circuit Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_virtual_circuit`





## Schema

### Required

- **connection_id** (String) UUID of Connection where the VC is scoped to
- **nni_vlan** (Number) Equinix Metal network-to-network VLAN ID
- **port_id** (String) UUID of the Connection Port where the VC is scoped to
- **project_id** (String) UUID of the Project where the VC is scoped to
- **vlan_id** (String) UUID of the VLAN to associate

### Optional

- **id** (String) The ID of this resource.
- **name** (String) Name of the Virtual Circuit resource

### Read-only

- **nni_vnid** (Number) Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
- **status** (String) Status of the virtual circuit resource
- **vnid** (Number) VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/


