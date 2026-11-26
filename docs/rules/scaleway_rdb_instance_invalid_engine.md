# scaleway_rdb_instance_invalid_engine

Disallow an invalid engine on `scaleway_rdb_instance`.

## Example

```hcl
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "postgres-16" // invalid engine!
}
```

```
$ tflint
1 issue(s) found:

Error: "postgres-16" is an invalid Managed Database engine. It must look like "PostgreSQL-16" or "MySQL-8". (scaleway_rdb_instance_invalid_engine)

  on main.tf line 4:
   4:   engine    = "postgres-16" // invalid engine!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_rdb_instance_invalid_engine.md
```

## Why

The provider does not validate `engine` at plan time. A wrong engine name (for instance `postgres-16`, borrowed from another cloud provider) is only rejected by the API when terraform apply runs, and changing `engine` on an existing Instance triggers a major upgrade workflow.

## How To Fix

Use `<engine>-<major version>` where engine is `PostgreSQL` or `MySQL` (case-insensitive), for example `PostgreSQL-16` or `MySQL-8`. The available versions are listed by `scw rdb engine list`.

This rule only checks the format and the engine name, not the version itself, so that it does not reject versions released after this ruleset.
