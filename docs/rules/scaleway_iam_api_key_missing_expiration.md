# scaleway_iam_api_key_missing_expiration

Require an expiration date on `scaleway_iam_api_key`.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_iam_api_key_missing_expiration" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_iam_api_key" "ci" {
  application_id = scaleway_iam_application.ci.id
}
```

```
$ tflint
1 issue(s) found:

Warning: The API key never expires. Set expires_at and rotate the key before that date. (scaleway_iam_api_key_missing_expiration)

  on main.tf line 1:
   1: resource "scaleway_iam_api_key" "ci" {

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_iam_api_key_missing_expiration.md
```

## Why

The Scaleway Landing Zone Framework requires API keys to be generated with strict validity limits and rotated regularly, to limit the impact of a leaked credential. Without `expires_at`, the key is valid forever.

## How To Fix

Set `expires_at` (RFC 3339 date). It cannot be changed afterwards, so rotate the key by creating a new one, for example with the `time_rotating` resource of the `hashicorp/time` provider.
