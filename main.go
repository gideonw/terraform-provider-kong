package main

import (
	"github.com/gideonw/terraform-provider-kong/kong"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kong.Provider})
}
