---
page_title: "metal_virtual_circuit Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_virtual_circuit`





## Schema

### Required

- **virtual_circuit_id** (String) ID of the virtual circuit to lookup

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **name** (String) Name of the virtual circuit
- **nni_vlan** (Number) Nni VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
- **nni_vnid** (Number) Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
- **project_id** (String) ID of the projct to which the virtual circuit belongs
- **status** (String) Status of the virtual circuit
- **vnid** (Number) VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/


