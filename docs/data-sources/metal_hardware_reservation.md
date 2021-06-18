---
page_title: "metal_hardware_reservation Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_hardware_reservation`





## Schema

### Optional

- **device_id** (String) UUID of device occupying the reservation
- **id** (String) ID of the hardware reservation to look up

### Read-only

- **facility** (String) Plan type for the reservation
- **plan** (String) Plan type for the reservation
- **project_id** (String) UUID of project this reservation is scoped to
- **provisionable** (Boolean) Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first
- **short_id** (String) Reservation short ID
- **spare** (Boolean) Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix
- **switch_uuid** (String) Switch short ID, can be used to determine if two devices are connected to the same switch


