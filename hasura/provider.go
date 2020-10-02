package hasura

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	graphql "github.com/hasura/go-graphql-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	defaultAPIEndpoint = "https://data.pro.hasura.io/v1/graphql"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultAPIEndpoint,
				DefaultFunc: schema.EnvDefaultFunc("HASURA_API_ENDPOINT", nil),
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access token for Hasura API",
				DefaultFunc: schema.EnvDefaultFunc("HASURA_ACCESS_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hasura_tenant": resourceTenant(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	var endpoint, accessToken string

	if endpoint = d.Get("api_endpoint").(string); endpoint == "" {
		return nil, diag.FromErr(fmt.Errorf("empty api_endpoint for Hasura"))
	}
	if accessToken = d.Get("access_token").(string); accessToken == "" {
		return nil, diag.FromErr(fmt.Errorf("empty access_token for Hasura"))
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient(endpoint, httpClient)
	return client, diags
}
