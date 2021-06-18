terraform {
  required_providers {
    metal = {
      source = "equinix/metal"
      version = "2.0.1"
    }
  }
}

resource "metal_reserved_ip_block" "example" {
  project_id = local.project_id
  facility   = "sv15"
  quantity   = 2
}

resource "metal_device" "example" {
  project_id       = local.project_id
  facilities       = ["sv15"]
  plan             = "c3.medium.x86"
  operating_system = "ubuntu_20_04"
  hostname         = "test"
  billing_cycle    = "hourly"

  ip_address {
    type            = "public_ipv4"
    cidr            = 31
    reservation_ids = [metal_reserved_ip_block.example.id]
  }

  ip_address {
    type = "private_ipv4"
  }
}
