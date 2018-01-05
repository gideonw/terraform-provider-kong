package kong

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/gideonw/gokong"
)

func resourceKongAPI() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongAPICreate,
		Read:   resourceKongAPIRead,
		Delete: resourceKongAPIDelete,
		Update: resourceKongAPIUpdate,
		Exists: resourceKongAPIExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uris": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"methods": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"upstream_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"strip_uri": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
			},
			"preserve_host": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},
			"retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Default:  5,
			},
			"upstream_connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Default:  60000,
			},
			"upstream_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Default:  60000,
			},
			"upstream_read_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Default:  60000,
			},
			"https_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"http_if_terminated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
			},
		},
	}
}

func resourceKongAPICreate(d *schema.ResourceData, meta interface{}) error {

	apiRequest := createKongAPIRequestFromResourceData(d)

	api, err := meta.(*gokong.KongAdminClient).APIs().CreateOrUpdate(apiRequest)

	if err != nil {
		return fmt.Errorf("failed to create or update kong api: %v error: %v", apiRequest, err)
	}

	d.SetId(api.ID)

	return resourceKongAPIRead(d, meta)
}

func resourceKongAPIUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)
	apiRequest := createKongAPIRequestFromResourceData(d)

	_, err := meta.(*gokong.KongAdminClient).APIs().UpdateByID(d.Id(), apiRequest)

	if err != nil {
		return fmt.Errorf("error updating kong api: %s", err)
	}

	return resourceKongAPIRead(d, meta)
}

func resourceKongAPIRead(d *schema.ResourceData, meta interface{}) error {

	api, err := meta.(*gokong.KongAdminClient).APIs().GetByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not find kong api: %v", err)
	}

	d.Set("name", api.Name)
	d.Set("hosts", api.Hosts)
	d.Set("uris", api.URIs)
	d.Set("methods", api.Methods)
	d.Set("upstream_url", api.UpstreamURL)
	d.Set("strip_uri", api.StripURI)
	d.Set("preserve_host", api.PreserveHost)
	d.Set("retries", api.Retries)
	d.Set("upstream_connect_timeout", api.UpstreamConnectTimeout)
	d.Set("upstream_send_timeout", api.UpstreamSendTimeout)
	d.Set("upstream_read_timeout", api.UpstreamReadTimeout)
	d.Set("https_only", api.HTTPSOnly)
	d.Set("http_if_terminated", api.HTTPIfTerminated)

	return nil
}

func resourceKongAPIDelete(d *schema.ResourceData, meta interface{}) error {

	err := meta.(*gokong.KongAdminClient).APIs().DeleteByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not delete kong api: %v", err)
	}

	return nil
}

func resourceKongAPIExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	apiName := d.Get("name").(string)

	api, err := meta.(*gokong.KongAdminClient).APIs().GetByName(apiName)
	if err != nil {
		return false, fmt.Errorf("could not find kong api: %v", err)
	}

	return api != nil && api.Name == apiName, nil
}

func createKongAPIRequestFromResourceData(d *schema.ResourceData) *gokong.APIRequest {

	apiRequest := &gokong.APIRequest{}

	apiRequest.Name = readStringFromResource(d, "name")
	apiRequest.Hosts = readStringArrayFromResource(d, "hosts")
	apiRequest.URIs = readStringArrayFromResource(d, "uris")
	apiRequest.Methods = readStringArrayFromResource(d, "methods")
	apiRequest.UpstreamURL = readStringFromResource(d, "upstream_url")
	apiRequest.StripURI = readBoolFromResource(d, "strip_uri")
	apiRequest.PreserveHost = readBoolFromResource(d, "preserve_host")
	apiRequest.Retries = readIntFromResource(d, "retries")
	apiRequest.UpstreamConnectTimeout = readIntFromResource(d, "upstream_connect_timeout")
	apiRequest.UpstreamSendTimeout = readIntFromResource(d, "upstream_send_timeout")
	apiRequest.UpstreamReadTimeout = readIntFromResource(d, "upstream_read_timeout")
	apiRequest.HTTPSOnly = readBoolFromResource(d, "https_only")
	apiRequest.HTTPIfTerminated = readBoolFromResource(d, "http_if_terminated")

	return apiRequest
}
