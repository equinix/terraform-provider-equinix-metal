terraform {
  required_providers {
    metal = {
      source = "equinix/metal"
    }
  }
}

resource "metal_project" "example" {
  name = "example"
}

resource "metal_vlan" "example" {
  project_id       = metal_project.example.id
  facility         = "sv15"
  description      = "example"
}
