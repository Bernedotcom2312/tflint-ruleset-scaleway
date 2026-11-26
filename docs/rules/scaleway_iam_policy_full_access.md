# scaleway_iam_policy_full_access

Disallow the `AllProductsFullAccess` permission set in `scaleway_iam_policy`.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_iam_policy_full_access" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_iam_policy" "devs" {
  group_id = scaleway_iam_group.devs.id

  rule {
    project_ids          = [scaleway_account_project.dev.id]
    permission_set_names = ["AllProductsFullAccess"]
  }
}
```

```
$ tflint
1 issue(s) found:

Warning: AllProductsFullAccess grants every permission on every product. Use product-specific permission sets, such as InstancesFullAccess. (scaleway_iam_policy_full_access)

  on main.tf line 6:
   6:     permission_set_names = ["AllProductsFullAccess"]

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_iam_policy_full_access.md
```

## Why

The Scaleway Landing Zone Framework applies the least privilege principle: use highly specific permission sets (for example `InstancesFullAccess` on a development Project) rather than `AllProductsFullAccess`, to limit the blast radius of a compromised account.

## How To Fix

Replace `AllProductsFullAccess` with the permission sets the principal actually needs (see `scw iam permission-set list`), and scope the rule to Projects with `project_ids` rather than to the whole Organization whenever possible.
