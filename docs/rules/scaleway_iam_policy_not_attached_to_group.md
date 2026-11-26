# scaleway_iam_policy_not_attached_to_group

Disallow `scaleway_iam_policy` attached directly to a user or an application.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_iam_policy_not_attached_to_group" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_iam_policy" "alice" {
  user_id = scaleway_iam_user.alice.id

  rule {
    project_ids          = [scaleway_account_project.dev.id]
    permission_set_names = ["InstancesFullAccess"]
  }
}
```

```
$ tflint
1 issue(s) found:

Warning: The policy is attached directly to a user. Attach it to an IAM group (group_id) and add the user to the group. (scaleway_iam_policy_not_attached_to_group)

  on main.tf line 2:
   2:   user_id = scaleway_iam_user.alice.id

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_iam_policy_not_attached_to_group.md
```

## Why

The Scaleway Landing Zone Framework recommends a strict RBAC model: attach IAM policies to IAM groups based on enterprise roles, and avoid assigning permissions directly to individual users or applications, which creates management overhead and security blind spots.

## How To Fix

Create an IAM group for the role, attach the policy to the group with `group_id`, and add the users or applications to the group:

```hcl
resource "scaleway_iam_group" "developers" {
  name     = "developers"
  user_ids = [scaleway_iam_user.alice.id]
}

resource "scaleway_iam_policy" "developers" {
  group_id = scaleway_iam_group.developers.id

  rule {
    project_ids          = [scaleway_account_project.dev.id]
    permission_set_names = ["InstancesFullAccess"]
  }
}
```
