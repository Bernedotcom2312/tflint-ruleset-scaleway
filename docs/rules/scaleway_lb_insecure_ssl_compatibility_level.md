# scaleway_lb_insecure_ssl_compatibility_level

Disallow Load Balancers accepting TLS versions older than 1.2.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_lb_insecure_ssl_compatibility_level" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_lb" "main" {
  type                    = "LB-S"
  ssl_compatibility_level = "ssl_compatibility_level_old"
}
```

```
$ tflint
1 issue(s) found:

Warning: ssl_compatibility_level_old accepts TLS versions older than 1.2. Use ssl_compatibility_level_intermediate (the default) or ssl_compatibility_level_modern. (scaleway_lb_insecure_ssl_compatibility_level)

  on main.tf line 3:
   3:   ssl_compatibility_level = "ssl_compatibility_level_old"

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_lb_insecure_ssl_compatibility_level.md
```

## Why

The Scaleway Landing Zone Framework asks to restrict the protocols accepted by Load Balancers to TLS 1.2 and TLS 1.3. `ssl_compatibility_level_old` also accepts deprecated TLS versions.

## How To Fix

Remove `ssl_compatibility_level` to use the default (`ssl_compatibility_level_intermediate`), or set `ssl_compatibility_level_modern` if every client supports TLS 1.3.
