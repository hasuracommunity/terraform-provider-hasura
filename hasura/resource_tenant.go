package hasura

import (
	"context"
	"log"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/mutation"

	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/query"

	"github.com/machinebox/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTenant() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Tenant",
		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:         schema.TypeString,
				Description:  "Cloud Provider",
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"aws", "gcp", "azure"}, false),
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
		CreateContext: resourceTenantCreate,
		ReadContext:   resourceTenantRead,
		UpdateContext: resourceTenantUpdate,
		DeleteContext: resourceTenantDelete,
	}
}

func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*graphql.Client)
	req := mutation.CreateTenant
	req.Var("name", d.Get("name").(string))
	req.Var("region", d.Get("region").(string))
	req.Var("cloud", d.Get("cloud").(string))
	req.Var("databaseUrl", d.Get("database_url").(string))

	var resp mutation.CreateTenantResponse
	if err := client.Run(ctx, req, &resp); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("Created Tenant with id:%s", resp.CreateTenant.ID)

	d.SetId(resp.CreateTenant.ID)
	return resourceTenantRead(ctx, d, meta)
}

func resourceTenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if diags := resourceTenantDelete(ctx, d, meta); diags.HasError() {
		return diags
	}

	return resourceTenantCreate(ctx, d, meta)
}

func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*graphql.Client)
	var diags diag.Diagnostics
	id := d.Id()
	var resp mutation.DeleteTenantResponse
	req := mutation.DeleteTenant
	req.Var("tenantId", id)
	if err := client.Run(ctx, req, &resp); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("Delete Tenant with id:%s", id)
	d.SetId("")
	return diags
}
