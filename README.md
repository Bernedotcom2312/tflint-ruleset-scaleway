# TFLint Ruleset for terraform-provider-scaleway

[![build](https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/actions/workflows/build.yml/badge.svg)](https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/actions/workflows/build.yml)
[![e2e](https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/actions/workflows/e2e.yml/badge.svg)](https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/actions/workflows/e2e.yml)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

[TFLint](https://github.com/terraform-linters/tflint) ruleset plugin for the [Terraform Scaleway Provider](https://registry.terraform.io/providers/scaleway/scaleway/latest/docs), inspired by [tflint-ruleset-aws](https://github.com/terraform-linters/tflint-ruleset-aws).

The Scaleway provider leaves many values unchecked until the API rejects them: an Instance commercial type, a Kubernetes pool node type, a Managed Database node type or engine, a Load Balancer type, a bucket name... `terraform plan` succeeds and `terraform apply` fails, sometimes halfway through a replacement. This ruleset catches those mistakes statically, in seconds, in your editor or your CI.

## Requirements

- TFLint v0.46+
- Go v1.27 (to build from source)

## Installation

Declare the plugin in `.tflint.hcl` and run `tflint --init`:

```hcl
plugin "scaleway" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/Bernedotcom2312/tflint-ruleset-scaleway"
}
```

## Rules

See the [list of rules](docs/rules/README.md). Rules checking values rejected by the API are enabled by default; best practice rules must be enabled explicitly:

```hcl
rule "scaleway_instance_security_group_inbound_default_policy" {
  enabled = true
}
```

Any rule can be disabled for a single resource with an annotation:

```hcl
resource "scaleway_rdb_instance" "preview" {
  # tflint-ignore: scaleway_rdb_instance_backup_disabled
  disable_backup = true
}
```

## Scaleway Landing Zone Framework

Best practice rules follow the guidelines of the Scaleway Landing Zone Framework: default-deny security groups, no administration port exposed to the internet, databases in Private Networks with encryption at rest and backups, TLS 1.2+, bucket versioning, IAM least privilege and tagging. See [the coverage of the framework](docs/landing-zone.md) and the configuration that enables all of them.

## How it works

Rules only read your configuration: no Scaleway credentials are needed to run them. Values that cannot be resolved statically (unknown variables, resource attributes) are skipped.

The commercial types and node types accepted by the rules live in [`rules/models/catalog.go`](rules/models/catalog.go). It is generated from the Scaleway API by [`tools/catalog`](tools/catalog), merging every zone and region so that a type available in a single locality is still accepted.

## Building the plugin

Clone the repository locally and run:

```console
$ make
```

Install the plugin in `~/.tflint.d/plugins` so that TFLint uses your local build:

```console
$ make install
```

With a local build, declare the plugin without `version` and `source`:

```hcl
plugin "scaleway" {
  enabled = true
}
```

Other targets:

| Target | Description |
| --- | --- |
| `make test` | Run the unit tests |
| `make e2e` | Run the end-to-end tests against the installed plugin (requires `tflint`) |
| `make vet` | Run `go vet` |
| `make fmt` | Format the code |
| `make catalog` | Refresh `rules/models/catalog.go` from the Scaleway API (requires `SCW_SECRET_KEY`) |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Mozilla Public License 2.0](LICENSE)
