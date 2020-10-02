package query

import "github.com/machinebox/graphql"

var GetTenantDetails = graphql.NewRequest(`
  query ($id: uuid!) {
       tenant_by_pk (id:$id) {
          id
          region
          cloud
      }
  }
`)

type GetTenantResponse struct {
	TenantByPK struct {
		ID     string
		Region string
		Cloud  string
	} `json:"tenant_by_pk"`
}
