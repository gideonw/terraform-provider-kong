package kong

import (
	"fmt"

	"github.com/gideonw/gokong"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKongPlugin() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongPluginCreate,
		Read:   resourceKongPluginRead,
		Delete: resourceKongPluginDelete,
		Update: resourceKongPluginUpdate,
		Exists: resourceKongPluginExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"consumer_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"config": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
				Default:  nil,
			},
		},
	}
}

func resourceKongPluginCreate(d *schema.ResourceData, meta interface{}) error {

	pluginRequest := createKongPluginRequestFromResourceData(d)

	plugin, err := meta.(*gokong.KongAdminClient).Plugins().UpdateOrAdd(pluginRequest)

	if err != nil {
		return fmt.Errorf("failed to create kong plugin: %v error: %v", pluginRequest, err)
	}

	d.SetId(plugin.ID)

	return resourceKongPluginRead(d, meta)
}

func resourceKongPluginUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	pluginRequest := createKongPluginRequestFromResourceData(d)

	_, err := meta.(*gokong.KongAdminClient).Plugins().UpdateByID(d.Id(), pluginRequest)

	if err != nil {
		return fmt.Errorf("error updating kong plugin: %s", err)
	}

	return resourceKongPluginRead(d, meta)
}

func resourceKongPluginRead(d *schema.ResourceData, meta interface{}) error {

	plugin, err := meta.(*gokong.KongAdminClient).Plugins().GetByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not find kong plugin: %v", err)
	}

	if plugin == nil {
		return fmt.Errorf("kong plugin nil: %s", d.Id())
	}

	d.Set("name", plugin.Name)

	return nil
}

func resourceKongPluginDelete(d *schema.ResourceData, meta interface{}) error {

	err := meta.(*gokong.KongAdminClient).Plugins().DeleteByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not delete kong plugin: %v", err)
	}

	return nil
}

func resourceKongPluginExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	pluginName := d.Get("name").(string)
	apiID := d.Get("api_id").(string)

	plugins, err := meta.(*gokong.KongAdminClient).Plugins().ListFiltered(&gokong.PluginFilter{
		APIID: apiID,
		Name:  pluginName,
	})
	if err != nil {
		return false, fmt.Errorf("could not find kong plugin: %v", err)
	}

	for _, v := range plugins.Results {
		if v != nil && v.Name == pluginName {
			return true, nil
		}
	}

	return false, nil
}

func createKongPluginRequestFromResourceData(d *schema.ResourceData) *gokong.PluginRequest {

	pluginRequest := &gokong.PluginRequest{}

	pluginRequest.Name = readStringFromResource(d, "name")
	pluginRequest.APIID = readStringFromResource(d, "api_id")
	pluginRequest.ConsumerID = readStringFromResource(d, "consumer_id")
	pluginRequest.Config = readMapFromResource(d, "config")

	return pluginRequest
}
