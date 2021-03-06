---
page_title: "Equinix Metal: metal_virtual_circuit"
subcategory: ""
description: |-
  Retrieve Equinix Fabric Virtual Circuit
---

# metal_virtual_circuit

Use this data source to retrieve a virtual circuit resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

## Example Usage

```hcl
data "metal_connection" "example_connection" {
  connection_id     = "4347e805-eb46-4699-9eb9-5c116e6a017d"
}

data "metal_virtual_circuit" "example_vc" {
  virtual_circuit_id = data.metal_connection.example_connection.ports[1].virtual_circuit_ids[0]
}

```

## Argument Reference

* `virtual_circuit_id` - (Required) ID of the virtual circuit resource

## Attributes Reference

* `name` - Name of the virtual circuit resource
* `status` - Status of the virtal circuit
* `project_id` - ID of project to which the VC belongs
* `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)

