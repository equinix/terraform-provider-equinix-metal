---
page_title: "metal_project_ssh_key Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_project_ssh_key`





## Schema

### Required

- **project_id** (String) The Equinix Metal project id of the Equinix Metal SSH Key

### Optional

- **id** (String) The id of the SSH Key
- **search** (String) The name, fingerprint, id, or public_key of the SSH Key to search for in the Equinix Metal project

### Read-only

- **created** (String)
- **fingerprint** (String)
- **name** (String) The label of the Equinix Metal SSH Key
- **owner_id** (String)
- **public_key** (String) The public SSH key that will be authorized for SSH access on Equinix Metal devices provisioned with this key
- **updated** (String)


