# scaleway_resource_missing_tags

Require tags on every Scaleway resource that supports them.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner", "cost-center"]

  # Optional: resource types to skip.
  exclude = ["scaleway_k8s_pool"]
}
```

Scaleway tags are free-form strings. A required key `env` is satisfied by any of the tags `env`, `env=<value>` or `env:<value>`. Object Storage buckets use S3-compatible map tags: the key must be present in the map.

## Example

```hcl
resource "scaleway_k8s_cluster" "main" {
  name = "prd-payments-kapsule"
  tags = ["env=prod", "owner=payments"]
}
```

```
$ tflint
1 issue(s) found:

Notice: The resource is missing the following tags: "cost-center". (scaleway_resource_missing_tags)

  on main.tf line 3:
   3:   tags = ["env=prod", "owner=payments"]

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_resource_missing_tags.md
```

## Why

The Scaleway Landing Zone Framework makes tagging an essential control: Scaleway resources are scoped to Projects, and tags are one of the primary mechanisms to rebuild cross-Project groupings for cost allocation, ownership tracking and operational visibility. As Scaleway has no native policy engine to enforce tagging standards, the framework asks to validate them in CI/CD pipelines.

Only resources whose schema has a `tags` argument are checked (see `rules/models/tags.go`). Tags that cannot be resolved statically are skipped.

## How To Fix

Add the missing tags, ideally from a shared `locals` block so that every module applies the same standard:

```hcl
locals {
  tags = ["env=${var.env}", "owner=${var.owner}", "cost-center=${var.cost_center}"]
}

resource "scaleway_k8s_cluster" "main" {
  name = "prd-payments-kapsule"
  tags = local.tags
}
```
