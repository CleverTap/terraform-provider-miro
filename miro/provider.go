package miro

import (
	"fmt"
	"terraform-provider-miro/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"miro_token": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				DefaultFunc: schema.EnvDefaultFunc("MIRO_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"miro_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"miro_user": dataSourceUser(),
		},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	miroToken 	:= d.Get("miro_token").(string)
	if len(miroToken) == 0 {
		return client.NewClient(miroToken), fmt.Errorf("Token not provided.")
	}
	return client.NewClient(miroToken), nil
}
