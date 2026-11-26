# scaleway_k8s_pool_invalid_node_type

Disallow an invalid node type on `scaleway_k8s_pool`.

## Example

```hcl
resource "scaleway_k8s_pool" "main" {
  cluster_id = scaleway_k8s_cluster.main.id
  name       = "default"
  size       = 3
  node_type  = "t2.micro" // invalid node type!
}
```

```
$ tflint
1 issue(s) found:

Error: "t2.micro" is an invalid Kubernetes pool node type. (scaleway_k8s_pool_invalid_node_type)

  on main.tf line 5:
   5:   node_type  = "t2.micro" // invalid node type!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_k8s_pool_invalid_node_type.md
```

## Why

The provider does not validate `node_type` at plan time: `terraform plan` succeeds and `terraform apply` fails when the pool is created. Because `node_type` forces a new pool, a typo introduced later also fails only at apply time, after the plan announced a replacement.

## How To Fix

Use an Instance commercial type (for example `PRO2-S` or `pro2_s`; case and `-`/`_` are ignored like the provider does), or `external` for Kosmos pools managed outside of Scaleway. See `scw k8s node-type list`.
