# hasura_tenant Resource

Provides a Hasura tenant resource.

## Example Usage

```hcl
# Create a new Spinnaker application
resource "hasura_tenant" "my_tenant" {
    cloud         = "aws"
    region        = "us-east-2"
    database_url  = "postgres://ec2-3-95-87-221.compute-1.amazonaws.com:5432/mydb"
}
```

## Argument Reference

The following arguments are supported.

* `cloud` - (Required) Name of the cloud provider of the database.
* `region` - (Required) Region of the database.
* `database_url` - (Required) URL of the database.
* `name` - (Optional) Name of the tenant. 

## Import

Applications can be imported using their Hasura tenant ID, e.g.

```console
$ terraform import hasura_tenant.my_tenant 4025c885-55ac-49b6-a961-114579fb88e3
```
