package metal

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/packethost/packngo"
)

func dataSourceMetalVirtualCircuit() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMetalVirtualCircuitRead,

		Schema: map[string]*schema.Schema{
			"virtual_circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the virtual circuit to lookup",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the virtual circuit",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the virtual circuit",
			},
			"vnid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/",
			},
			"nni_vnid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/",
			},
			"nni_vlan": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Nni VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/",
			},
			"project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the projct to which the virtual circuit belongs",
			},
		},
	}
}

func dataSourceMetalVirtualCircuitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*packngo.Client)
	vcId := d.Get("virtual_circuit_id").(string)

	vc, _, err := client.VirtualCircuits.Get(
		vcId,
		&packngo.GetOptions{Includes: []string{"project"}})
	if err != nil {
		return err
	}

	d.SetId(vc.ID)
	return setMap(d, map[string]interface{}{
		"virtual_circuit_id": vc.ID,
		"name":               vc.Name,
		"status":             vc.Status,
		"vnid":               vc.VNID,
		"nni_vnid":           vc.NniVNID,
		"nni_vlan":           vc.NniVLAN,
		"project_id":         vc.Project.ID,
	})
}
