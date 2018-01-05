package kong

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/gideonw/gokong"
)

func dataSourceKongAPI() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKongAPIRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"upstream_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uris": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"methods": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"upstream_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"strip_uri": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"preserve_host": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"retries": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upstream_connect_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upstream_send_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upstream_read_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"https_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"http_if_terminated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceKongAPIRead(d *schema.ResourceData, meta interface{}) error {

	filter := &gokong.APIFilter{}

	if v, _ := d.GetOk("filter"); v != nil {
		filterSet := v.(*schema.Set).List()
		if len(filterSet) == 1 {
			filterMap := filterSet[0].(map[string]interface{})
			filter.ID = filterMap["id"].(string)
			filter.Name = filterMap["name"].(string)
			filter.UpstreamURL = filterMap["upstream_url"].(string)
		}
	}

	results, err := meta.(*gokong.KongAdminClient).APIs().ListFiltered(filter)

	if err != nil {
		return fmt.Errorf("could not find api, error: %v", err)
	}

	if len(results.Results) == 0 {
		return fmt.Errorf("could not find api using filter: %v", filter)
	}

	if len(results.Results) > 1 {
		return fmt.Errorf("found more than 1 api make filter more restrictive")
	}

	api := results.Results[0]

	d.SetId(api.ID)
	d.Set("id", api.ID)
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
