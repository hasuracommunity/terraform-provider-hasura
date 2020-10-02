package mutation

import (
	"github.com/hasura/go-graphql-client"
)

var CreateProject struct {
	CreateTenant struct {
		ID graphql.ID
	} `graphql:"createTenant(cloud: $cloud, region: $region, databaseUrl: $databaseUrl, name: $name)"`
}

var DeleteTenant struct {
	DeleteTenant struct {
		Status graphql.String
	} `graphql:"deleteTenant(tenantId: $tenantId)"`
}
