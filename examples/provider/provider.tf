terraform {
  required_providers {
    metal = {
      source = "equinix/metal"
      # version = "1.0.0"
    }
  }
}

# Configure the Equinix Metal Provider.
provider "metal" {
  auth_token = var.auth_token
}

data "metal_project" "project" {
  name = "My Project"
}

# If you want to create a fresh project, you can create one with metal_project
#
# resource "metal_project" "cool_project" {
#   name           = "My First Terraform Project"
# }

# Create a device and add it to tf_project_1
resource "metal_device" "web1" {
  hostname         = "web1"
  plan             = "c3.medium.x86"
  facilities       = ["ny5"]
  operating_system = "ubuntu_20_04"
  billing_cycle    = "hourly"
  project_id       = data.metal_project.project.id

  # if you created a project with the metal_project resource, refer to its ID
  # project_id       = metal_project.cool_project.id

  # You can find the ID of your project in the URL of the Equinix Metal console.
  # For example, if you see your devices listed at
  # https://console.equinix.com/projects/352000fb2-ee46-4673-93a8-de2c2bdba33b
  # .. then 352000fb2-ee46-4673-93a8-de2c2bdba33b is your project ID.
}
