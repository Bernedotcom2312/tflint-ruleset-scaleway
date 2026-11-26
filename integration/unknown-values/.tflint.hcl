plugin "terraform" {
  enabled = false
}

plugin "scaleway" {
  enabled = true
}

rule "scaleway_instance_security_group_admin_port_exposed" {
  enabled = true
}

rule "scaleway_rdb_instance_public_endpoint" {
  enabled = true
}

rule "scaleway_redis_cluster_public_endpoint" {
  enabled = true
}

rule "scaleway_rdb_instance_encryption_at_rest_disabled" {
  enabled = true
}

rule "scaleway_lb_insecure_ssl_compatibility_level" {
  enabled = true
}

rule "scaleway_object_bucket_versioning_disabled" {
  enabled = true
}

rule "scaleway_iam_api_key_missing_expiration" {
  enabled = true
}

rule "scaleway_iam_policy_full_access" {
  enabled = true
}

rule "scaleway_iam_policy_not_attached_to_group" {
  enabled = true
}

rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}

rule "scaleway_instance_security_group_inbound_default_policy" {
  enabled = true
}

rule "scaleway_rdb_instance_backup_disabled" {
  enabled = true
}
