package hasura

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hasuracommunity/terraform-provider-hasura/hasura/graphql/query"
	"github.com/machinebox/graphql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSourceHasuraTenant_basic(t *testing.T) {
	resourceName := "hasura_tenant.test"
	rDBURL := acctest.RandomWithPrefix("postgres://")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHasuraTenantCheckDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccHasuraTenant_basic(rDBURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasuraTenantExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "cloud", "aws"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-2"),
					resource.TestCheckResourceAttr(resourceName, "database_url", rDBURL),
				),
			},
		},
	})
}

func testAccCheckHasuraTenantExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("tenant not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no tentant ID is set")
		}
		client := testAccProvider.Meta().(graphql.Client)
		err := resource.Retry(1*time.Minute, func() *resource.RetryError {
			req := query.GetTenantDetails
			req.Var("id", n)

			var resp query.GetTenantResponse

			if err := client.Run(context.Background(), req, &resp); err != nil {
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("unable to find tenant after retries: %s", err)
		}

		return nil
	}
}

func testAccCheckHasuraTenantCheckDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("hasura Tenant not found, tenant: %s", n)
		}

		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("no ID is set")
		}

		req := query.GetTenantDetails
		req.Var("id", id)

		client := testAccProvider.Meta().(graphql.Client)
		var resp query.GetTenantResponse

		retry := 5
		for {
			if err := client.Run(context.Background(), req, &resp); err != nil {
				return fmt.Errorf("could not get tenant, error:%v", err)
			}

			if resp.TenantByPK.ID == "" {
				return nil
			}

			if resp.TenantByPK.ID != "" {
				retry--
				log.Printf("[INFO] Retring CheckDestroy in 1 seconds, retry count: %v", 5-retry)
				time.Sleep(1 * time.Second)
			}

			if retry <= 0 {
				break
			}
		}

		return fmt.Errorf("hasura tenant still exists, ID: %s", resp.TenantByPK.ID)
	}
}

func testAccHasuraTenant_basic(rDBURL string) string {
	return fmt.Sprintf(`
resource "hasura_tenant" "test" {
	cloud        = "aws"
    region       = "us-east-2"
	database_url = %q
}
`, rDBURL)
}
