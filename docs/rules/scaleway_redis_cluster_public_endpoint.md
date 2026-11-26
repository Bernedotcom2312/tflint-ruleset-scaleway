# scaleway_redis_cluster_public_endpoint

Disallow Managed Redis clusters exposed on a public endpoint.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_redis_cluster_public_endpoint" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = "RED1-micro"
  user_name = "admin"
  password  = var.redis_password
}
```

```
$ tflint
1 issue(s) found:

Warning: The Redis cluster has no private_network block, so it is exposed on a public endpoint. Attach it to a Private Network. (scaleway_redis_cluster_public_endpoint)

  on main.tf line 1:
   1: resource "scaleway_redis_cluster" "main" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_redis_cluster_public_endpoint.md
```

## Why

The Scaleway Landing Zone Framework recommends deploying databases exclusively within Private Networks, without any public IP address. Without a `private_network` block, the provider exposes the cluster on a public network.

## How To Fix

Attach the cluster to a Private Network:

```hcl
resource "scaleway_redis_cluster" "main" {
  # ...

  private_network {
    id = scaleway_vpc_private_network.data.id
  }
}
```

If a public endpoint is required, restrict it with `acl` blocks and ignore the rule for that resource with a `tflint-ignore` annotation.
