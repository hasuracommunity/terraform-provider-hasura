package hasura

import (
	"context"
	"log"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/query"

	"github.com/machinebox/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTenant() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Tenant",
		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:        schema.TypeString,
				Description: "Cloud Provider",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the backend",
				Computed:    true,
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the backend",
				Required:    true,
			},
			"database_url": {
				Type:        schema.TypeString,
				Description: "URL of the backend",
				Sensitive:   true,
				Required:    true,
			},
		},
		ReadContext: dataSourceRead,
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*graphql.Client)
	var diags diag.Diagnostics
	id := d.Id()
	req := query.GetTenantDetails
	req.Var("id", id)

	var resp query.GetTenantResponse
	if err := client.Run(ctx, req, &resp); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("Get Tenant with id:%s,cloud:%s,region:%s", resp.TenantByPK.ID, resp.TenantByPK.Cloud, resp.TenantByPK.Region)

	if v := resp.TenantByPK.Region; v != "" {
		d.Set("region", v)
	}

	if v := resp.TenantByPK.Cloud; v != "" {
		d.Set("cloud", v)
	}

	return diags
}
