# scaleway_rdb_instance_backup_disabled

Disallow disabling automated backups on `scaleway_rdb_instance`.

This rule is disabled by default.

## Configuration

```hcl
rule "scaleway_rdb_instance_backup_disabled" {
  enabled = true
}
```

## Example

```hcl
resource "scaleway_rdb_instance" "main" {
  name           = "main"
  node_type      = "DB-DEV-S"
  engine         = "PostgreSQL-16"
  disable_backup = true
}
```

```
$ tflint
1 issue(s) found:

Warning: Automated backups are disabled for this Database Instance. (scaleway_rdb_instance_backup_disabled)

  on main.tf line 5:
   5:   disable_backup = true

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_rdb_instance_backup_disabled.md
```

## Why

Without automated backups, a deleted table, a failed migration or a deleted Instance cannot be restored from a recent point in time. The Scaleway Landing Zone Framework requires database backups to be enabled at creation time, with no possibility of omission, at least daily.

## How To Fix

Remove `disable_backup` (backups are enabled by default) or set it to `false`, and tune `backup_schedule_frequency` and `backup_schedule_retention` if needed. If backups are intentionally disabled (ephemeral environments), disable the rule for that resource with an annotation:

```hcl
resource "scaleway_rdb_instance" "preview" {
  # tflint-ignore: scaleway_rdb_instance_backup_disabled
  disable_backup = true
}
```
