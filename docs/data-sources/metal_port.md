---
page_title: "metal_port Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_port`





## Schema

### Optional

- **device_id** (String) Device UUID where to lookup the port
- **id** (String) UUID of the port to lookup
- **name** (String) Name of the port to look up, e.g. bond0, eth1

### Read-only

- **bond_id** (String) UUID of the bond port
- **bond_name** (String) Name of the bond port
- **bonded** (Boolean) Flag indicating whether the port is bonded
- **disbond_supported** (Boolean) Flag indicating whether the port can be removed from a bond
- **mac** (String) MAC address of the port
- **native_vlan_id** (String) UUID of native VLAN of the port
- **network_type** (String) One of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
- **type** (String) Port type
- **vlan_ids** (List of String) UUIDs of attached VLANs


