# scaleway_rdb_instance_invalid_node_type

Disallow an invalid node type on `scaleway_rdb_instance`.

## Example

```hcl
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  engine    = "PostgreSQL-16"
  node_type = "db-dev-xxs" // invalid node type!
}
```

```
$ tflint
1 issue(s) found:

Error: "db-dev-xxs" is an invalid Managed Database node type. (scaleway_rdb_instance_invalid_node_type)

  on main.tf line 4:
   4:   node_type = "db-dev-xxs" // invalid node type!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_rdb_instance_invalid_node_type.md
```

## Why

The provider does not validate `node_type` at plan time: `terraform plan` succeeds and `terraform apply` fails when the Database Instance is created or resized.

## How To Fix

Pick a node type listed by `scw rdb node-type list` or on the [Managed Databases pricing page](https://www.scaleway.com/en/pricing/managed-databases/). The comparison is case-insensitive, like the provider.
