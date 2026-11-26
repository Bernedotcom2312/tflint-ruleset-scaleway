# scaleway_rdb_instance_public_endpoint

Disallow Managed Database Instances exposed on a public endpoint.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_rdb_instance_public_endpoint" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"
}
```

```
$ tflint
1 issue(s) found:

Warning: The Database Instance has no private_network block, so it gets a public endpoint by default. Attach it to a Private Network. (scaleway_rdb_instance_public_endpoint)

  on main.tf line 1:
   1: resource "scaleway_rdb_instance" "main" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_rdb_instance_public_endpoint.md
```

## Why

The Scaleway Landing Zone Framework recommends deploying databases exclusively within Private Networks, without any public IP address.

The provider creates a public endpoint when the `load_balancer` block is set, but also by default when neither `private_network` nor `load_balancer` is set. The rule reports both cases.

## How To Fix

Attach the Database Instance to a Private Network and remove the `load_balancer` block:

```hcl
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"

  private_network {
    pn_id       = scaleway_vpc_private_network.data.id
    enable_ipam = true
  }
}
```

If a public endpoint is required, restrict it with `scaleway_rdb_acl` and ignore the rule for that resource with a `tflint-ignore` annotation.
