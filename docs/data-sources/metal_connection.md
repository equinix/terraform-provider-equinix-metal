---
page_title: "metal_connection Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_connection`





## Schema

### Required

- **connection_id** (String) ID of the connection to lookup

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **description** (String) Description of the connection resource
- **facility** (String) Slug of a facility to which the connection belongs
- **metro** (String) Slug of a metro to which the connection belongs
- **name** (String) Name of the connection resource
- **organization_id** (String) ID of organization to which the connection belongs
- **ports** (List of Object) List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`) (see [below for nested schema](#nestedatt--ports))
- **redundancy** (String) Connection redundancy - reduntant or primary
- **speed** (Number) Port speed in bits per second
- **status** (String) Status of the connection resource
- **token** (String) Fabric Token for the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
- **type** (String) Connection type - dedicated or shared

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


