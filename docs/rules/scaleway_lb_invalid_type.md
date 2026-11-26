# scaleway_lb_invalid_type

Disallow an invalid commercial type on `scaleway_lb`.

## Example

```hcl
resource "scaleway_lb" "main" {
  name = "main"
  type = "LB-XXL" // invalid type!
}
```

```
$ tflint
1 issue(s) found:

Error: "LB-XXL" is an invalid Load Balancer commercial type. (scaleway_lb_invalid_type)

  on main.tf line 3:
   3:   type = "LB-XXL" // invalid type!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_lb_invalid_type.md
```

## Why

The provider does not validate `type` at plan time: `terraform plan` succeeds and `terraform apply` fails when the Load Balancer is created or migrated.

## How To Fix

Pick a commercial type listed by `scw lb lb-types list` or on the [Network pricing page](https://www.scaleway.com/en/pricing/network/) (`LB-S`, `LB-GP-M`, `LB-GP-L`, `LB-GP-XL`). The comparison is case-insensitive, like the provider.
