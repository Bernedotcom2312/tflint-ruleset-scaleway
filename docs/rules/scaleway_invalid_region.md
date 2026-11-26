# scaleway_invalid_region

Disallow unknown regions on Scaleway resources, data sources and on the `scaleway` provider block.

## Example

```hcl
resource "scaleway_object_bucket" "main" {
  name   = "my-bucket"
  region = "fr-par-1" // this is a zone, not a region!
}
```

```
$ tflint
1 issue(s) found:

Warning: "fr-par-1" is an invalid Scaleway region. (scaleway_invalid_region)

  on main.tf line 3:
   3:   region = "fr-par-1" // this is a zone, not a region!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_invalid_region.md
```

## Why

The provider only emits a warning for an unknown region, which is easy to miss in a long plan output. The API then rejects the request at `terraform apply`. Mixing up a region (`fr-par`) and a zone (`fr-par-1`), or reusing a region name from another cloud provider, are common mistakes.

## How To Fix

Use one of the known regions: `fr-par`, `nl-ams`, `pl-waw`, `it-mil`.

The severity is `warning`, like the provider, because a newly opened region might not be known by this ruleset yet. Please open an issue if that happens.
