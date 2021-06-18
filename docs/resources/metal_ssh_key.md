---
page_title: "metal_ssh_key Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_ssh_key`





## Schema

### Required

- **name** (String) The name of the SSH key for identification
- **public_key** (String) The public key. If this is a file, it

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **created** (String) The timestamp for when the SSH key was created
- **fingerprint** (String) The fingerprint of the SSH key
- **owner_id** (String) The UUID of the Equinix Metal API User who owns this key
- **updated** (String) The timestamp for the last time the SSH key was updated


