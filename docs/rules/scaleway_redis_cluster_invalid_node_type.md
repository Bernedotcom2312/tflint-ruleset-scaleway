# scaleway_redis_cluster_invalid_node_type

Disallow an invalid node type on `scaleway_redis_cluster`.

## Example

```hcl
resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = "RED1-XXL" // invalid node type!
}
```

```
$ tflint
1 issue(s) found:

Error: "RED1-XXL" is an invalid Managed Redis node type. (scaleway_redis_cluster_invalid_node_type)

  on main.tf line 4:
   4:   node_type = "RED1-XXL" // invalid node type!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_redis_cluster_invalid_node_type.md
```

## Why

The provider does not validate `node_type` at plan time: `terraform plan` succeeds and `terraform apply` fails when the cluster is created.

## How To Fix

Pick a node type listed by `scw redis node-type list`. The comparison is case-insensitive, like the provider.
