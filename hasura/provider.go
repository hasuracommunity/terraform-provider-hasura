package hasura

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	graphql "github.com/machinebox/graphql"

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
				DefaultFunc: schema.EnvDefaultFunc("HASURA_API_ENDPOINT", defaultAPIEndpoint),
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
		DataSourcesMap: map[string]*schema.Resource{
			"hasura_tenant": dataSourceTenant(),
		},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

type AddAccessTokenTransport struct {
	T           http.RoundTripper
	AccessToken string
}

func (adt *AddAccessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("pat %s", adt.AccessToken))
	return adt.T.RoundTrip(req)
}

func newAccessTokenTransport(accessToken string) *AddAccessTokenTransport {
	t := http.DefaultTransport
	return &AddAccessTokenTransport{t, accessToken}
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

	httpClient := &http.Client{
		Transport: newAccessTokenTransport(accessToken),
	}

	opts := graphql.WithHTTPClient(httpClient)
	client := graphql.NewClient(endpoint, opts)
	return client, diags
}
