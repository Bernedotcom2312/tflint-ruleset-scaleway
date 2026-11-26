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

### Release verification

Releases are verified by `tflint --init` with [GitHub artifact attestations](https://docs.github.com/en/actions/security-for-github-actions/using-artifact-attestations/using-artifact-attestations-to-establish-provenance-for-builds) by default. `checksums.txt` is also signed with the following PGP key (fingerprint `A8B5 2113 1E4A D989 3E87 5C5A 80D4 D468 4F2D EEC9`), which you can require instead:

```hcl
plugin "scaleway" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/Bernedotcom2312/tflint-ruleset-scaleway"

  signature   = "pgp"
  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  mQINBGsIerYBEADNGz/wa31LIDkLM7DJn9Jj5SFH8iCY1U0FAYktl1mctR1hTA5D
  Ywh4As00V8ii/gC6hGk6eeCjEoAc6UGa7nEUvHkOpUtwVXM51ZpPkyMjzF59a8IT
  TTZE9CQOby/QYTDIKDirSMnCDHSpBYzikgI8FCSzv9q2Dy7zTD+8baqGnv2Pnv/3
  DiEDNZS0n35pstHdvekm4GVEs2stPrCTqLg+q7o4OC3S9Yp+JH7k4ULtVY32l71x
  VE+CkuehFl7ntySmRTqcnbnTCCejoofGMKOMdafmf7CR2TncLuz/1AVoyHiR0A/Y
  fRGV8T4G+7QmVezB7kFpbOr4eTdLxcFwpO5jxvIK/30XJYh9EauJwuPXsN+j6da4
  WEArEzHP9Q4nn+bOY7PNigbeg/F+9seEy8z+6r5xshggs4eOzVpLu+2zUyp8qv+L
  hDfQPy3x79ESp24jSKIcGvIkCggpLWQ7yOZPWTIsJTWz5Ex1LeE+hVRsx4NCEe8v
  dlstdAxfPi2e4O4xsC8KuEEY5trOw/totbYIdVASydpi5my8x2prKTQKeK/d+qN7
  GfEjkVtVEgPF8yxSJY8V3c7P9CVUY1ILd7nClXHXhiP1D3IUg7A6SMAwox433QJa
  uArIUp/DynBildOL3gthSEY0GMi0XdCElCgyuc+RgIf98WraKfJPL1TXiwARAQAB
  tCpWUFkgQ2xvdWQgQ29uc3VsdGluZyA8Y29udGFjdEB2cHljbG91ZC5mcj6JAnME
  EwEIAF0WIQSotSETHkrZiT6HXFqA1NRoTy3uyQUCawh6thsUgAAAAAAEAA5tYW51
  MiwyLjUrMS4xMiwwLDMCGwMFCQPCZwAFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcC
  F4AACgkQgNTUaE8t7slz5xAAzBlYjahC/YKmBaOdCu3en0O694EaVoUm9O+A9dsZ
  N1+ca1m0o9tbbJ1EB9mNU0aq4Q+80bmpvjxvpwi6XNx9xjMhkvtoOYd6iFNajNOp
  vgTS4kUSW4r2ROuMVHLCJi0CwXJBJI3egr4zRtXz2eyXR1aQU+Pgu71bvLJ5+cjs
  vV+DtNjlaV/jntApIxG7ZBfUb2VvwSMytwH3d0O9JQKjri6gK6wQO4d1Ju0Q7BZj
  F/JC+1aNcxd7dtfnu28wOUmm2gzf6X35rMXvEIuTuOl/n+85Z2VviiSF6aEX8pRt
  /cwzDzaVMQ3QTGkIusQJv2AXcohA9RzJOhdT4zj/Ir4mojwKiN0K8YTJMyxKwy2h
  Yl0AbpGCZ0NlKMEMHJnoouDSPO/XU1SOcj60E6br3S2hY2VCQMg1Jzp9Uzn6AR6C
  OS1G6Ds1ffy6trlmZcdbtZp6ptQWj8f9FON8LMNI3yuBxn2m2rKJx2lP0/ORfUx+
  IgdCxxd03cQ8u15HtePHDbg+LW9Kbxe1KEooiPuGqn5mJR8rYZ1Gza9LfgT7GdsU
  AaDsUdRYtY4z1XSnYmvGvN8yVy0D9rSTaW+GWLu8KUhRioDSu1jgDuKFjcUnPopg
  g/uliDvAd0HpuY2oUTwOcMwXGGT5tmyXpmXPX568GpzFP0NrGRzVoPGUJuMIN25x
  VFE=
  =QdCN
  -----END PGP PUBLIC KEY BLOCK-----
  KEY
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
