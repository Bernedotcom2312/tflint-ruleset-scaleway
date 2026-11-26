# scaleway_object_bucket_versioning_disabled

Require versioning on `scaleway_object_bucket`.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_object_bucket_versioning_disabled" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_object_bucket" "backups" {
  name = "acme-backups"
}
```

```
$ tflint
1 issue(s) found:

Warning: Versioning is not enabled on this bucket. Add a versioning block with enabled = true. (scaleway_object_bucket_versioning_disabled)

  on main.tf line 1:
   1: resource "scaleway_object_bucket" "backups" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_object_bucket_versioning_disabled.md
```

## Why

The Scaleway Landing Zone Framework recommends enabling bucket versioning to protect against accidental deletion or overwrites, in particular for backup repositories and Terraform state buckets.

## How To Fix

```hcl
resource "scaleway_object_bucket" "backups" {
  name = "acme-backups"

  versioning {
    enabled = true
  }
}
```

Once enabled, versioning can only be suspended, never removed. Combine it with `lifecycle_rule` blocks to expire old versions. For buckets that do not need versioning (caches, public assets), use a `tflint-ignore` annotation.
