package metal

var (
	resourceDescriptions = map[string]string{
		"volume": "Provides an Equinix Metal Block Storage Volume resource to allow you to manage block volumes on your account.  Once created by Terraform, they must then be attached and mounted using the api and `metal_block_attach` and `metal_block_detach` scripts.",
	}
	schemaDescriptions = map[string]string{
		"facility":      "The facility in which the resource should be created",
		"project_id":    "ID of metap project in which the resource should be created",
		"locked":        "Setting this parameter will prevent resource deletion",
		"billing_cycle": "hourly (default) or monthly",
		"state":         "Provisioning state of the resource",
	}
)
