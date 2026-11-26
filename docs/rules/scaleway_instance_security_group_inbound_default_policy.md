# scaleway_instance_security_group_inbound_default_policy

Disallow security groups that accept inbound traffic by default.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_instance_security_group_inbound_default_policy" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_instance_security_group" "web" {
  name = "web"

  inbound_rule {
    action = "accept"
    port   = 443
  }
}
```

```
$ tflint
1 issue(s) found:

Warning: inbound_default_policy is not set and defaults to "accept". Set it to "drop" and allow traffic with explicit inbound rules. (scaleway_instance_security_group_inbound_default_policy)

  on main.tf line 1:
   1: resource "scaleway_instance_security_group" "web" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_instance_security_group_inbound_default_policy.md
```

## Why

The Scaleway Landing Zone Framework asks for a default-deny policy on all inbound traffic. While security groups created in the console drop inbound traffic by default, the Terraform provider defaults `inbound_default_policy` to `accept`: every port that no `inbound_rule` drops is reachable. In the example above, the `accept` rule on port 443 has no effect and SSH, databases or any other listening service stay exposed.

## How To Fix

Set `inbound_default_policy = "drop"` and allow the expected traffic with explicit `inbound_rule` blocks:

```hcl
resource "scaleway_instance_security_group" "web" {
  name                   = "web"
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 443
  }
}
```
