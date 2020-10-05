# Hasura Provider

The [Hasura](https://github.com/hasura) is a GraphQL Engine gives you fast, instant realtime GraphQL on any Postgres application, existing or new. 
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Hasura provider
provider "hasura" {
  access_token = "${var.access_token}"
}
```

## Argument Reference

The following arguments are supported:

* `api_endpoint` - (Optional) Endpoint for the Hasura API. Defaults to `https://data.pro.hasura.io/v1/graphql`. You can use `HASURA_ENDPOINT` environment variable too.
* `access_token` - (Required) Access token to access the API. You can use `HASURA_ACCESS_TOKEN` environment variable too.
