locals {
  tags = ["env=prod", "owner=platform"]
}

resource "scaleway_vpc_private_network" "data" {
  name = "prd-shared-pn-data"
  tags = local.tags
}

resource "scaleway_instance_security_group" "bastion" {
  inbound_default_policy = "drop"
  tags                   = ["env=prod"]

  inbound_rule {
    action = "accept"
    port   = 22
  }
}

resource "scaleway_rdb_instance" "main" {
  name      = "prd-payments-db"
  node_type = "DB-PRO2-XXS"
  engine    = "PostgreSQL-16"
  tags      = local.tags
}

resource "scaleway_rdb_instance" "compliant" {
  name               = "prd-billing-db"
  node_type          = "DB-PRO2-XXS"
  engine             = "PostgreSQL-16"
  encryption_at_rest = true
  tags               = local.tags

  private_network {
    pn_id       = scaleway_vpc_private_network.data.id
    enable_ipam = true
  }
}

resource "scaleway_redis_cluster" "cache" {
  name      = "prd-payments-cache"
  version   = "7.2.11"
  node_type = "RED1-micro"
  user_name = "admin"
  password  = var.redis_password
  tags      = local.tags
}

resource "scaleway_lb" "public" {
  name                    = "prd-payments-lb"
  type                    = "LB-S"
  ssl_compatibility_level = "ssl_compatibility_level_old"
  tags                    = local.tags
}

resource "scaleway_object_bucket" "state" {
  name = "acme-terraform-state"
  tags = {
    env   = "prod"
    owner = "platform"
  }
}

resource "scaleway_iam_application" "ci" {
  name = "ci"
  tags = local.tags
}

resource "scaleway_iam_api_key" "ci" {
  application_id = scaleway_iam_application.ci.id
}

resource "scaleway_iam_policy" "ci" {
  application_id = scaleway_iam_application.ci.id
  tags           = local.tags

  rule {
    project_ids          = ["11111111-1111-1111-1111-111111111111"]
    permission_set_names = ["AllProductsFullAccess"]
  }
}

variable "redis_password" {
  type      = string
  sensitive = true
}
