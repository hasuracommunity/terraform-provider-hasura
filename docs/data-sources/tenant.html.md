# hasura_tenant Data Source

Provides a Hasura tenant data source.

## Example Usage

```hcl
# Create a new Spinnaker application
data "hasura_tenant" "my_tenant" {}
```

## Argument Reference

The following arguments are supported.

* `id` - ID of the Hasura tenant.
* `cloud` - Name of the cloud provider of the database.
* `region` - Region of the database.
* `database_url` - URL of the database.
* `name` - Name of the tenant.

## Import

Applications can be imported using their Hasura tenant ID, e.g.

```console
$ terraform import hasura_tenant.my_tenant 4025c885-55ac-49b6-a961-114579fb88e3
```
