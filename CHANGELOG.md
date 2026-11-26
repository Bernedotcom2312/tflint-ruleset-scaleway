# Changelog

All notable changes to this project are documented in this file. The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-11-26

### Added

- Possible error rules, enabled by default:
  - `scaleway_instance_server_invalid_type`
  - `scaleway_k8s_pool_invalid_node_type`
  - `scaleway_k8s_cluster_invalid_cni`
  - `scaleway_rdb_instance_invalid_node_type`
  - `scaleway_rdb_instance_invalid_engine`
  - `scaleway_redis_cluster_invalid_node_type`
  - `scaleway_lb_invalid_type`
  - `scaleway_object_bucket_invalid_name`
- Locality rules, enabled by default: `scaleway_invalid_zone`, `scaleway_invalid_region`.
- Best practice rules following the Scaleway Landing Zone Framework, disabled by default:
  - `scaleway_instance_security_group_inbound_default_policy`
  - `scaleway_instance_security_group_admin_port_exposed`
  - `scaleway_rdb_instance_public_endpoint`
  - `scaleway_redis_cluster_public_endpoint`
  - `scaleway_rdb_instance_encryption_at_rest_disabled`
  - `scaleway_rdb_instance_backup_disabled`
  - `scaleway_lb_insecure_ssl_compatibility_level`
  - `scaleway_object_bucket_versioning_disabled`
  - `scaleway_iam_api_key_missing_expiration`
  - `scaleway_iam_policy_full_access`
  - `scaleway_iam_policy_not_attached_to_group`
  - `scaleway_resource_missing_tags`
- Coverage of the Landing Zone Framework in `docs/landing-zone.md`.
- `tools/catalog` to refresh the Scaleway catalog from the API.

[Unreleased]: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/releases/tag/v0.1.0
