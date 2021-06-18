---
page_title: "metal_project Resource - tpm"
subcategory: ""
description: |-
  
---

# Resource `metal_project`





## Schema

### Required

- **name** (String) The name of the project

### Optional

- **backend_transfer** (Boolean) Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
- **bgp_config** (Block List, Max: 1) Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/) (see [below for nested schema](#nestedblock--bgp_config))
- **id** (String) The ID of this resource.
- **organization_id** (String) The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account
- **payment_method_id** (String) The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with organization_id, or default)

### Read-only

- **created** (String) The timestamp for when the project was created
- **updated** (String) The timestamp for the last time the project was updated

<a id="nestedblock--bgp_config"></a>
### Nested Schema for `bgp_config`

Required:

- **asn** (Number) Autonomous System Number for local BGP deployment
- **deployment_type** (String) "local" or "global", the local is likely to be usable immediately, the global will need to be review by Equinix Metal engineers

Optional:

- **md5** (String, Sensitive) Password for BGP session in plaintext (not a checksum)

Read-only:

- **max_prefix** (Number) The maximum number of route filters allowed per server
- **status** (String) Status of BGP configuration in the project


