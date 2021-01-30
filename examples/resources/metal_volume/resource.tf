locals {
    project_id = "552345b2-ee46-4673-93a8-de2c2bdba33b"
}


resource "metal_volume" "volume1" {
  description   = "terraform-volume-1"
  facility      = "ewr1"
  project_id    = local.project_id
  plan          = "storage_1"
  size          = 100
  billing_cycle = "hourly"

  snapshot_policies {
    snapshot_frequency = "1day"
    snapshot_count     = 7
  }

  snapshot_policies {
    snapshot_frequency = "1month"
    snapshot_count     = 6
  }
}
