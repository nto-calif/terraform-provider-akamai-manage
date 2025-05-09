package main

import (
	"github.com/nto-calif/terraform-provider-akamai-manage/akamai_manage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: akamai_manage.Provider,
	})
}
