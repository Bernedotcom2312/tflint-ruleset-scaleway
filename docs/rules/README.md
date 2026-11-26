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
<!-- errors -->

## Localities

These rules check the zones and regions used by Scaleway resources, data sources and provider blocks.

|Name|Severity|Enabled|
| --- | --- | --- |
|[scaleway_invalid_zone](scaleway_invalid_zone.md)|Warning|✔|
|[scaleway_invalid_region](scaleway_invalid_region.md)|Warning|✔|
<!-- localities -->
