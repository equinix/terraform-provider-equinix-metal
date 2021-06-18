---
page_title: "metal_device Resource - tpm"
subcategory: ""
description: |-
  This resource can be used to create, modify, and delete devices in Equinix Metal.
---

# Resource `metal_device`

This resource can be used to create, modify, and delete devices in Equinix Metal.

## Example Usage

```terraform
# Create a device

resource "metal_device" "web1" {
  hostname         = "tf.coreos2"
  plan             = "c3.small.x86"
  metro            = "sv"
  operating_system = "ubuntu_20_04"
  billing_cycle    = "hourly"
  project_id       = local.project_id
}


# Create device via iPXE boot, using the Ignition Provider for provisioning

resource "metal_device" "pxe1" {
  hostname         = "tf.coreos2-pxe"
  plan             = "c3.small.x86"
  metro            = "sv"
  operating_system = "custom_ipxe"
  billing_cycle    = "hourly"
  project_id       = local.project_id
  ipxe_script_url  = "https://rawgit.com/cloudnativelabs/pxe/master/metal/coreos-stable-metal.ipxe"
  always_pxe       = "false"
  user_data        = data.ignition_config.example.rendered
}


# Create a device without public IP address in facility ny5, with only a /30 private IPv4 subnet (4 IP addresses)

resource "metal_device" "web1" {
  hostname         = "tf.coreos2"
  plan             = "c3.small.x86"
  facilities       = ["ny5"]
  operating_system = "ubuntu_20_04"
  billing_cycle    = "hourly"
  project_id       = local.project_id
  ip_address {
    type = "private_ipv4"
    cidr = 30
  }
}


# Deploy device on next-available reserved hardware and do custom partitioning.

resource "metal_device" "web1" {
  hostname                = "tftest"
  plan                    = "c3.small.x86"
  facilities              = ["ny5"]
  operating_system        = "ubuntu_20_04"
  billing_cycle           = "hourly"
  project_id              = local.project_id
  hardware_reservation_id = "next-available"
  storage                 = <<EOS
{
  "disks": [
    {
      "device": "/dev/sda",
      "wipeTable": true,
      "partitions": [
        {
          "label": "BIOS",
          "number": 1,
          "size": "4096"
        },
        {
          "label": "SWAP",
          "number": 2,
          "size": "3993600"
        },
        {
          "label": "ROOT",
          "number": 3,
          "size": "0"
        }
      ]
    }
  ],
  "filesystems": [
    {
      "mount": {
        "device": "/dev/sda3",
        "format": "ext4",
        "point": "/",
        "create": {
          "options": [
            "-L",
            "ROOT"
          ]
        }
      }
    },
    {
      "mount": {
        "device": "/dev/sda2",
        "format": "swap",
        "point": "none",
        "create": {
          "options": [
            "-L",
            "SWAP"
          ]
        }
      }
    }
  ]
}
EOS
}
```

## Schema

### Required

- **billing_cycle** (String) monthly or hourly
- **hostname** (String) The device name
- **operating_system** (String) The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response
- **plan** (String) The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response
- **project_id** (String) The ID of the project in which to create the device

### Optional

- **always_pxe** (Boolean) If true, a device with OS custom_ipxe will
- **custom_data** (String, Sensitive) A string of the desired Custom Data for the device
- **description** (String) Description string for the device
- **facilities** (List of String) List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or any (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with metro
- **force_detach_volumes** (Boolean) Delete device even if it has volumes attached. Only applies for destroy action
- **hardware_reservation_id** (String) The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your next available reservation automatically
- **id** (String) The ID of this resource.
- **ip_address** (Block List) A list of IP address types for the device (structure is documented below) (see [below for nested schema](#nestedblock--ip_address))
- **ipxe_script_url** (String) URL pointing to a hosted iPXE script. More
- **metro** (String) Metro area for the new device. Conflicts with facilities
- **project_ssh_key_ids** (List of String) Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the [metal_project_ssh_key](project_ssh_key.md) resource
- **storage** (String) JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc
- **tags** (List of String) Tags attached to the device
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **user_data** (String, Sensitive) A string of the desired User Data for the device
- **wait_for_reservation_deprovision** (Boolean) Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019)

### Read-only

- **access_private_ipv4** (String) The ipv4 private IP assigned to the device
- **access_public_ipv4** (String) The ipv4 maintenance IP assigned to the device
- **access_public_ipv6** (String) The ipv6 maintenance IP assigned to the device
- **created** (String) The timestamp for when the device was created
- **deployed_facility** (String) The facility where the device is deployed
- **deployed_hardware_reservation_id** (String) ID of hardware reservation where this device was deployed. It is useful when using the next-available hardware reservation
- **locked** (Boolean) Whether the device is locked
- **network** (List of Object) The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 addresses: public ipv4, private ipv4 and ipv6 (see [below for nested schema](#nestedatt--network))
- **network_type** (String, Deprecated) Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
- **ports** (List of Object) Ports assigned to the device (see [below for nested schema](#nestedatt--ports))
- **root_password** (String, Sensitive) Root password to the server (disabled after 24 hours)
- **ssh_key_ids** (List of String) List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
- **state** (String) The status of the device
- **updated** (String) The timestamp for the last time the device was updated

<a id="nestedblock--ip_address"></a>
### Nested Schema for `ip_address`

Required:

- **type** (String) one of public_ipv4,private_ipv4,public_ipv6

Optional:

- **cidr** (Number) CIDR suffix for IP block assigned to this device
- **reservation_ids** (List of String) IDs of reservations to pick the blocks from


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)
- **delete** (String)
- **update** (String)


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


