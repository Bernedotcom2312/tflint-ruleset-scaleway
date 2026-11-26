# Contributing

Thanks for your interest in improving the ruleset!

## Commit messages

This repository follows [Conventional Commits](https://www.conventionalcommits.org/). Keep commits atomic: one logical change per commit, with its tests and documentation.

| Type | Use for |
| --- | --- |
| `feat` | A new rule or a user-facing capability, e.g. `feat(rules): add scaleway_lb_invalid_type` |
| `fix` | A false positive, a false negative or a crash |
| `docs` | Documentation only |
| `test` | Tests only |
| `refactor` | Code changes that do not change behavior |
| `build` | Go modules, Makefile, GoReleaser |
| `ci` | GitHub Actions workflows |
| `chore` | Anything else, e.g. a catalog refresh: `chore(models): refresh catalog` |

Pull requests are squash-merged and their title is checked against the convention.

## Adding a rule

1. Create `rules/<rule_name>.go` and `rules/<rule_name>_test.go`. Rule names start with `scaleway_` and follow the tflint-ruleset-aws naming (`<resource>_invalid_<attribute>` for values the API rejects).
2. Register the rule in [`rules/provider.go`](rules/provider.go).
3. Document it in `docs/rules/<rule_name>.md` and add it to [`docs/rules/README.md`](docs/rules/README.md).
4. Add a case to the end-to-end tests in [`integration`](integration) when relevant.
5. Run `make fmt vet test`, then `make install e2e`.

Guidelines:

- Rules that catch values rejected by the API at apply time are enabled by default with the `error` severity.
- Best practice rules are disabled by default with the `warning` severity.
- Prefer false negatives over false positives: when a list of accepted values may grow (versions, new localities), check the format or use the `warning` severity.

## Refreshing the catalog

```console
$ SCW_SECRET_KEY=<secret key> make catalog
```

Review the diff and commit it with `chore(models): refresh catalog`.

## Releasing

1. Bump `Version` in [`project/main.go`](project/main.go) and the version in the README, and update [CHANGELOG.md](CHANGELOG.md).
2. Commit with `chore(release): vX.Y.Z` and push a `vX.Y.Z` tag.

The release workflow builds the binaries with GoReleaser and signs `checksums.txt` with the GPG key stored in the `GPG_PRIVATE_KEY` and `PASSPHRASE` secrets, as required by `tflint --init`.
