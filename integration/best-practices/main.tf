resource "scaleway_instance_security_group" "default" {
  name = "default"
}

resource "scaleway_instance_security_group" "open" {
  name                   = "open"
  inbound_default_policy = "accept"
}

resource "scaleway_instance_security_group" "web" {
  name                   = "web"
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 443
  }
}

resource "scaleway_rdb_instance" "main" {
  name           = "main"
  node_type      = "DB-DEV-S"
  engine         = "PostgreSQL-16"
  disable_backup = true
}
