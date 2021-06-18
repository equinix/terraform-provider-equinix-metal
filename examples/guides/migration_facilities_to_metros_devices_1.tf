resource "metal_device" "node" {
  project_id       = local.project_id
  facilities       = ["sv15"]
  plan             = "c3.small.x86"
  operating_system = "ubuntu_16_04"
  hostname         = "test"
  billing_cycle    = "hourly"
}
