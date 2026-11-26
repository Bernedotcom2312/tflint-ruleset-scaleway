# Changelog

All notable changes to this project are documented in this file. The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
- Best practice rules, disabled by default: `scaleway_instance_security_group_inbound_default_policy`, `scaleway_rdb_instance_backup_disabled`.
- `tools/catalog` to refresh the Scaleway catalog from the API.
