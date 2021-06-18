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
