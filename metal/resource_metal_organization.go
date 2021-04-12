package metal

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/packethost/packngo"
)

func resourceMetalOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetalOrganizationCreate,
		Read:   resourceMetalOrganizationRead,
		Update: resourceMetalOrganizationUpdate,
		Delete: resourceMetalOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},

			"website": {
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},

			"twitter": {
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},

			"logo": {
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMetalOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*packngo.Client)

	createRequest := &packngo.OrganizationCreateRequest{
		Name: d.Get("name").(string),
	}

	if attr, ok := d.GetOk("website"); ok {
		createRequest.Website = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		createRequest.Description = attr.(string)
	}

	if attr, ok := d.GetOk("twitter"); ok {
		createRequest.Twitter = attr.(string)
	}

	if attr, ok := d.GetOk("logo"); ok {
		createRequest.Logo = attr.(string)
	}

	org, _, err := client.Organizations.Create(createRequest)
	if err != nil {
		return friendlyError(err)
	}

	d.SetId(org.ID)

	return resourceMetalOrganizationRead(d, meta)
}

func resourceMetalOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*packngo.Client)

	key, _, err := client.Organizations.Get(d.Id(), nil)
	if err != nil {
		err = friendlyError(err)

		// If the project somehow already destroyed, mark as succesfully gone.
		if isNotFound(err) {
			d.SetId("")

			return nil
		}

		return err
	}

	d.SetId(key.ID)

	return setMap(d, map[string]interface{}{
		"name":        key.Name,
		"description": key.Description,
		"website":     key.Website,
		"twitter":     key.Twitter,
		"logo":        key.Logo,
		"created":     key.Created,
		"updated":     key.Updated,
	})
}

func resourceMetalOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*packngo.Client)

	updateRequest := &packngo.OrganizationUpdateRequest{}

	if d.HasChange("name") {
		oName := d.Get("name").(string)
		updateRequest.Name = &oName
	}

	if d.HasChange("description") {
		oDescription := d.Get("description").(string)
		updateRequest.Description = &oDescription
	}

	if d.HasChange("website") {
		oWebsite := d.Get("website").(string)
		updateRequest.Website = &oWebsite
	}

	if d.HasChange("twitter") {
		oTwitter := d.Get("twitter").(string)
		updateRequest.Twitter = &oTwitter
	}

	if d.HasChange("logo") {
		oLogo := d.Get("logo").(string)
		updateRequest.Logo = &oLogo
	}
	_, _, err := client.Organizations.Update(d.Id(), updateRequest)
	if err != nil {
		return friendlyError(err)
	}

	return resourceMetalOrganizationRead(d, meta)
}

func resourceMetalOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*packngo.Client)

	resp, err := client.Organizations.Delete(d.Id())
	if ignoreResponseErrors(httpForbidden, httpNotFound)(resp, err) != nil {
		return friendlyError(err)
	}

	d.SetId("")
	return nil
}
