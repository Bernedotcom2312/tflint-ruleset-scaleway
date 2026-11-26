# scaleway_invalid_zone

Disallow unknown Availability Zones on Scaleway resources, data sources and on the `scaleway` provider block.

## Example

```hcl
provider "scaleway" {
  zone = "fr-par-1"
}

resource "scaleway_instance_server" "main" {
  type = "DEV1-S"
  zone = "fr-par-4" // unknown zone!
}
```

```
$ tflint
1 issue(s) found:

Warning: "fr-par-4" is an invalid Scaleway zone. (scaleway_invalid_zone)

  on main.tf line 7:
   7:   zone = "fr-par-4" // unknown zone!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_invalid_zone.md
```

## Why

The provider only emits a warning for an unknown zone, which is easy to miss in a long plan output. The API then rejects the request at `terraform apply`. Mixing up a zone (`fr-par-1`) and a region (`fr-par`) is a common mistake.

## How To Fix

Use one of the known zones: `fr-par-1`, `fr-par-2`, `fr-par-3`, `nl-ams-1`, `nl-ams-2`, `nl-ams-3`, `pl-waw-1`, `pl-waw-2`, `pl-waw-3`, `it-mil-1`, `it-mil-2`, `it-mil-3`.

The severity is `warning`, like the provider, because a newly opened zone might not be known by this ruleset yet. Please open an issue if that happens.
