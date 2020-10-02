package hasura

import (
	"context"

	"github.com/hasura/go-graphql-client"

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
	//id := d.Id()
	//q := query.GetTenantDetails
	//
	//vars := map[string]interface{}{
	//	"id": graphql.ID(id),
	//}
	//
	//if err := client.Query(ctx, &q, vars); err != nil {
	//	return diag.FromErr(err)
	//}
	//
	//if v := q.TenantByPK.Region; v != "" {
	//	d.Set("region", v)
	//}
	//
	//if v := q.TenantByPK.Cloud; v != "" {
	//	d.Set("cloud", v)
	//}
	//
	//if v, ok := q.TenantByPK.ID.(string); ok {
	//	d.Set("id", v)
	//}
	_ = client

	return diags
}
