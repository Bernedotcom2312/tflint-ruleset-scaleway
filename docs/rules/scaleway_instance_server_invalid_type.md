# scaleway_instance_server_invalid_type

Disallow an invalid commercial type on `scaleway_instance_server`.

## Example

```hcl
resource "scaleway_instance_server" "main" {
  image = "ubuntu_noble"
  type  = "DEV1-XXL" // invalid type!
}
```

```
$ tflint
1 issue(s) found:

Error: "DEV1-XXL" is an invalid Instance commercial type. (scaleway_instance_server_invalid_type)

  on main.tf line 3:
   3:   type  = "DEV1-XXL" // invalid type!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_instance_server_invalid_type.md
```

## Why

The provider does not validate `type` at plan time: `terraform plan` succeeds and `terraform apply` fails when the Instance is created.

## How To Fix

Pick a commercial type listed by `scw instance server-type list` or on the [Instances pricing page](https://www.scaleway.com/en/pricing/virtual-instances/). The comparison is case-insensitive, like the provider.
