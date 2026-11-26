# scaleway_rdb_instance_encryption_at_rest_disabled

Require encryption at rest on `scaleway_rdb_instance`.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_rdb_instance_encryption_at_rest_disabled" {
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

Warning: Encryption at rest is not enabled for this Database Instance. Set encryption_at_rest = true. (scaleway_rdb_instance_encryption_at_rest_disabled)

  on main.tf line 1:
   1: resource "scaleway_rdb_instance" "main" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_rdb_instance_encryption_at_rest_disabled.md
```

## Why

The Scaleway Landing Zone Framework recommends encrypting data at rest. Managed Databases for PostgreSQL and MySQL offer an optional encryption at rest (LUKS, keys managed by Scaleway), which is disabled unless `encryption_at_rest = true`. Once enabled, it cannot be disabled.

## How To Fix

Set `encryption_at_rest = true`.
