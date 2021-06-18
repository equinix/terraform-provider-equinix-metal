terraform {
  required_providers {
    packet = {
      source = "packethost/packet"
      version = "3.2.1"
    }
  }
}

resource "packet_reserved_ip_block" "example" {
  project_id = local.project_id
  facility   = "sv15"
  quantity   = 2
}

resource "packet_device" "example" {
  project_id       = local.project_id
  facilities       = ["sv15"]
  plan             = "c3.medium.x86"
  operating_system = "ubuntu_20_04"
  hostname         = "test"
  billing_cycle    = "hourly"

  ip_address {
    type            = "public_ipv4"
    cidr            = 31
    reservation_ids = [packet_reserved_ip_block.example.id]
  }

  ip_address {
    type = "private_ipv4"
  }
}
