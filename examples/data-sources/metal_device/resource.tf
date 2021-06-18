# Fetch a device data by hostname and show it's ID

data "metal_device" "test" {
  project_id = local.project_id
  hostname   = "mydevice"
}

output "id" {
  value = data.metal_device.test.id
}


# Fetch a device data by ID and show its public IPv4
data "metal_device" "test" {
    device_id = "eb5ab7a3-ecea-43bb-8007-253d49ec29ec"
}

output "ipv4" {
  value = data.metal_device.test.access_public_ipv4
}
