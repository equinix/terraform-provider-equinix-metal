---
page_title: "Equinix Metal: metal_facility"
subcategory: ""
description: |-
  Provides an Equinix Metal facility datasource. This can be used to read facilities.
---

# metal_facility

Provides an Equinix Metal facility datasource.

## Example Usage

```hcl
# Fetch a facility by code and show its ID

data "metal_facility" "ny5" {
    code = "ny5"
}

output "id" {
  value = data.metal_facility.ny5.id
}
```

## Argument Reference

The following arguments are supported:

* `code` - The facility code

Facilities can be looked up by `code`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the facility
* `name` - The name of the facility
* `features` - The features of the facility
* `metro` - The metro code the facility is part of
