package mutation

import "github.com/hasura/go-graphql-client"

var CreateProject struct {
	CreateTenant struct {
		ID graphql.String
	} `graphql:"createTenant(cloud: $cloud, region: $region, databaseUrl: $databaseUrl, name: $name)"`
}
