plugin "terraform" {
  enabled = false
}

plugin "scaleway" {
  enabled = true
}

rule "scaleway_instance_security_group_inbound_default_policy" {
  enabled = true
}

rule "scaleway_rdb_instance_backup_disabled" {
  enabled = true
}
