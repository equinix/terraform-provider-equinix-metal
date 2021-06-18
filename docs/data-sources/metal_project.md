---
page_title: "metal_project Data Source - tpm"
subcategory: ""
description: |-
  
---

# Data Source `metal_project`





## Schema

### Optional

- **id** (String) The ID of this resource.
- **name** (String) The name which is used to look up the project
- **project_id** (String) The UUID by which to look up the project

### Read-only

- **backend_transfer** (Boolean) Whether Backend Transfer is enabled for this project
- **bgp_config** (List of Object) Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/) (see [below for nested schema](#nestedatt--bgp_config))
- **created** (String) The timestamp for when the project was created
- **organization_id** (String) The UUID of this project's parent organization
- **payment_method_id** (String) The UUID of payment method for this project
- **updated** (String) The timestamp for the last time the project was updated
- **user_ids** (List of String) List of UUIDs of user accounts which belong to this project

<a id="nestedatt--bgp_config"></a>
### Nested Schema for `bgp_config`

Read-only:

- **asn** (Number)
- **deployment_type** (String)
- **max_prefix** (Number)
- **md5** (String)
- **status** (String)


