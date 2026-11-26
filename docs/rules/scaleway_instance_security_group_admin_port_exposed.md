# scaleway_instance_security_group_admin_port_exposed

Disallow security group rules that accept administration ports (SSH and RDP by default) from the whole internet.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_instance_security_group_admin_port_exposed" {
  enabled = true

  # Optional: ports to protect. Defaults to [22, 3389].
  ports = [22, 3389, 5432]
}
```

## Example

```hcl
resource "scaleway_instance_security_group" "web" {
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 22
  }
}
```

```
$ tflint
1 issue(s) found:

Warning: Inbound rule accepts administration port 22 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead. (scaleway_instance_security_group_admin_port_exposed)

  on main.tf line 4:
   4:   inbound_rule {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_instance_security_group_admin_port_exposed.md
```

## Why

The Scaleway Landing Zone Framework asks to never expose administration ports directly to the public internet, and to route administrative access through a bastion or a Public Gateway instead.

The rule checks `inbound_rule` blocks of `scaleway_instance_security_group` and `scaleway_instance_security_group_rules`, with the provider defaults: `protocol` defaults to `TCP`, a rule without `ip_range` (nor `ip`) applies to `0.0.0.0/0`, and a rule without `port` (nor `port_range`) applies to every port.

## How To Fix

Remove the rule and use the SSH bastion of a Public Gateway (`bastion_enabled = true` on `scaleway_vpc_public_gateway`), a VPN, or restrict `ip_range` to trusted networks.
