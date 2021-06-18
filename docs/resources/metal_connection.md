---
page_title: "metal_connection Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_connection`





## Schema

### Required

- **name** (String) Name of the connection resource
- **organization_id** (String) ID of the organization responsible for the connection
- **redundancy** (String) Connection redundancy - redundant or primary
- **type** (String) Connection type - dedicated or shared

### Optional

- **description** (String) Description of the connection resource
- **facility** (String) Facility where the connection will be created
- **id** (String) The ID of this resource.
- **metro** (String) Metro where to the connection will be created
- **project_id** (String) ID of the project where the connection is scoped to, only used for type == "shared"

### Read-only

- **ports** (List of Object) List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`) (see [below for nested schema](#nestedatt--ports))
- **speed** (Number) Port speed in bits per second
- **status** (String) Status of the connection resource
- **token** (String) Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)

<a id="nestedatt--ports"></a>
### Nested Schema for `ports`

Read-only:

- **id** (String)
- **link_status** (String)
- **name** (String)
- **role** (String)
- **speed** (Number)
- **status** (String)
- **virtual_circuit_ids** (List of String)


