package hasura

import (
	"context"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/query"

	"github.com/hasura/go-graphql-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTenant() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Tenant",
		Schema: map[string]*schema.Schema{
			"cloud": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Cloud Provider",
				Required:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the backend",
				Computed:    true,
				Optional:    true,
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Region of the backend",
				Required:    true,
			},
			"database_url": &schema.Schema{
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
	q := query.GetTenantDetails

	vars := map[string]interface{}{
		"id": graphql.ID(id),
	}

	if err := client.Query(ctx, &q, vars); err != nil {
		return diag.FromErr(err)
	}

	if v := q.TenantByPK.Region; v != "" {
		d.Set("region", v)
	}

	if v := q.TenantByPK.Cloud; v != "" {
		d.Set("cloud", v)
	}

	//if v, ok := q.TenantByPK.ID.(string); ok {
	//	d.Set("id", v)
	//}

	return diags
}
