package hasura

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(data *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
