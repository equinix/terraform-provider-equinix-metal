---
page_title: "metal_device Data Source - tpm"
subcategory: ""
description: |-
  This datasource can be used to read existing devices in Equinix Metal.
---

# Data Source `metal_device`

This datasource can be used to read existing devices in Equinix Metal.



## Schema

### Optional

- **device_id** (String) Device ID
- **hostname** (String) The device name
- **id** (String) The ID of this resource.
- **project_id** (String) The id of the project in which the devices exists

### Read-only

- **access_private_ipv4** (String) The ipv4 private IP assigned to the device
- **access_public_ipv4** (String) The ipv4 management IP assigned to the device
- **access_public_ipv6** (String) The ipv6 management IP assigned to the device
- **always_pxe** (Boolean)
- **billing_cycle** (String) The billing cycle of the device (monthly or hourly)
- **description** (String) Description string for the device
- **facility** (String) The facility where the device is deployed
- **hardware_reservation_id** (String) The id of hardware reservation which this device occupies
- **ipxe_script_url** (String)
- **metro** (String) The metro where the device is deployed
- **network** (List of Object) The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks: ublic IPv4 at metal_device.name.network.0, IPv6 at metal_device.name.network.1 and private IPv4 at metal_device.name.network.2. Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list). (see [below for nested schema](#nestedatt--network))
- **network_type** (String) L2 network type of the device, one oflayer3, hybrid, layer2-individual, layer2-bonded
- **operating_system** (String) The operating system running on the device
- **plan** (String) The hardware config of the device
- **ports** (List of Object) Ports assigned to the device (see [below for nested schema](#nestedatt--ports))
- **root_password** (String, Sensitive) Root password to the server (if still available)
- **ssh_key_ids** (List of String) List of IDs of SSH keys deployed in the device, can be both user or project SSH keys
- **state** (String) The state of the device
- **storage** (String)
- **tags** (List of String) Tags attached to the device

<a id="nestedatt--network"></a>
### Nested Schema for `network`

Read-only:

- **address** (String)
- **cidr** (Number)
- **family** (Number)
- **gateway** (String)
- **public** (Boolean)


<a id="nestedatt--ports"></a>
### Nested Schema for `ports`

Read-only:

- **bonded** (Boolean)
- **id** (String)
- **mac** (String)
- **name** (String)
- **type** (String)


