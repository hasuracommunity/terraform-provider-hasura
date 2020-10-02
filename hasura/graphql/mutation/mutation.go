package mutation

import "github.com/machinebox/graphql"

var CreateTenant = graphql.NewRequest(`
   mutation ($databaseUrl: String!, $region: String!, $cloud: String!, $name: String) {
       createTenant (name:$name, region:$region, databaseUrl:$databaseUrl, cloud:$cloud) {
           id
           name
       }
   }
`)

type CreateTenantResponse struct {
	CreateTenant struct {
		ID   string
		Name string
	}
}
