terraform {
  required_providers {
    packet = {
      source = "packethost/packet"
    }
  }
}

resource "packet_project" "example" {
  name = "example"
}

resource "packet_vlan" "example" {
  project_id       = packet_project.example.id
  facility         = "sv15"
  description      = "example"
}
