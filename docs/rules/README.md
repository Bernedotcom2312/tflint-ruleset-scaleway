# Rules

## Possible Errors

These rules warn of possible errors that can occur at `terraform apply`. Rules marked with ✔ are enabled by default.

|Name|Severity|Enabled|
| --- | --- | --- |
|[scaleway_instance_server_invalid_type](scaleway_instance_server_invalid_type.md)|Error|✔|
|[scaleway_k8s_pool_invalid_node_type](scaleway_k8s_pool_invalid_node_type.md)|Error|✔|
|[scaleway_rdb_instance_invalid_node_type](scaleway_rdb_instance_invalid_node_type.md)|Error|✔|
|[scaleway_redis_cluster_invalid_node_type](scaleway_redis_cluster_invalid_node_type.md)|Error|✔|
|[scaleway_lb_invalid_type](scaleway_lb_invalid_type.md)|Error|✔|
|[scaleway_k8s_cluster_invalid_cni](scaleway_k8s_cluster_invalid_cni.md)|Error|✔|
|[scaleway_rdb_instance_invalid_engine](scaleway_rdb_instance_invalid_engine.md)|Error|✔|
|[scaleway_object_bucket_invalid_name](scaleway_object_bucket_invalid_name.md)|Error|✔|

## Localities

These rules check the zones and regions used by Scaleway resources, data sources and provider blocks.

|Name|Severity|Enabled|
| --- | --- | --- |
|[scaleway_invalid_zone](scaleway_invalid_zone.md)|Warning|✔|
|[scaleway_invalid_region](scaleway_invalid_region.md)|Warning|✔|

## Best Practices

These rules enforce best practices. They are disabled by default and must be enabled in `.tflint.hcl`.

|Name|Severity|Enabled|
| --- | --- | --- |
|[scaleway_instance_security_group_inbound_default_policy](scaleway_instance_security_group_inbound_default_policy.md)|Warning||
|[scaleway_rdb_instance_backup_disabled](scaleway_rdb_instance_backup_disabled.md)|Warning||
|[scaleway_instance_security_group_admin_port_exposed](scaleway_instance_security_group_admin_port_exposed.md)|Warning||
|[scaleway_rdb_instance_public_endpoint](scaleway_rdb_instance_public_endpoint.md)|Warning||
|[scaleway_redis_cluster_public_endpoint](scaleway_redis_cluster_public_endpoint.md)|Warning||
|[scaleway_rdb_instance_encryption_at_rest_disabled](scaleway_rdb_instance_encryption_at_rest_disabled.md)|Warning||
|[scaleway_lb_insecure_ssl_compatibility_level](scaleway_lb_insecure_ssl_compatibility_level.md)|Warning||
|[scaleway_iam_api_key_missing_expiration](scaleway_iam_api_key_missing_expiration.md)|Warning||
|[scaleway_object_bucket_versioning_disabled](scaleway_object_bucket_versioning_disabled.md)|Warning||
|[scaleway_iam_policy_full_access](scaleway_iam_policy_full_access.md)|Warning||
