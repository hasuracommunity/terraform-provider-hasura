package query

import (
	"github.com/hasura/go-graphql-client"
)

var GetTenantDetails struct {
	TenantByPK struct {
		ID     graphql.ID
		Cloud  graphql.String
		Region graphql.String
	} `graphql:"tenant_by_pk(id: $id)"`
}
