package hasura

import (
	"context"
	"fmt"
	"log"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/query"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/mutation"

	"github.com/hasura/go-graphql-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTenant() *schema.Resource {
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
		CreateContext: resourceTenantCreate,
		ReadContext:   resourceTenantRead,
		UpdateContext: resourceTenantUpdate,
		DeleteContext: resourceTenantDelete,
	}
}

func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*graphql.Client)

	m := mutation.CreateProject
	vars := map[string]interface{}{
		"name":        graphql.String(d.Get("name").(string)),
		"cloud":       graphql.String(d.Get("cloud").(string)),
		"region":      graphql.String(d.Get("region").(string)),
		"databaseUrl": graphql.String(d.Get("database_url").(string)),
	}

	if err := client.Mutate(ctx, &m, vars); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("Created Tenant with id:%s", m.CreateTenant.ID)

	d.SetId(m.CreateTenant.ID.(string))
	return resourceTenantRead(ctx, d, meta)
}

func resourceTenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if diags := resourceTenantDelete(ctx, d, meta); diags.HasError() {
		return diag.FromErr(fmt.Errorf("failed to delete tenant"))
	}

	return resourceTenantCreate(ctx, d, meta)
}

func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*graphql.Client)
	var diags diag.Diagnostics
	id := d.Id()

	q := mutation.DeleteTenant
	vars := map[string]interface{}{
		"tenantId": graphql.ID(id),
	}

	if err := client.Mutate(ctx, &q, vars); err != nil {
		d.SetId("")
	}

	return diags
}
