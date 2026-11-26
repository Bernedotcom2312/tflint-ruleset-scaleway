# scaleway_object_bucket_invalid_name

Disallow invalid names on `scaleway_object_bucket`.

## Example

```hcl
resource "scaleway_object_bucket" "main" {
  name = "My_Bucket" // invalid name!
}
```

```
$ tflint
1 issue(s) found:

Error: "My_Bucket" is an invalid bucket name: it can only contain lowercase letters, numbers, dots (.) and hyphens (-). (scaleway_object_bucket_invalid_name)

  on main.tf line 2:
   2:   name = "My_Bucket" // invalid name!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_object_bucket_invalid_name.md
```

## Why

The provider does not validate bucket names: `terraform plan` succeeds and `terraform apply` fails when the bucket is created.

## How To Fix

Scaleway Object Storage is S3-compatible and follows the DNS-compatible bucket naming rules. A bucket name must:

- be between 3 and 63 characters long,
- only contain lowercase letters, numbers, dots (`.`) and hyphens (`-`),
- begin and end with a lowercase letter or a number,
- not be formatted as an IP address.

Names containing interpolations that cannot be resolved statically (for example `"logs-${random_id.suffix.hex}"`) are skipped.
