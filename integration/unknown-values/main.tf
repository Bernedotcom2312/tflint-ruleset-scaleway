# Values that cannot be resolved statically must never raise an issue.

variable "zone" {
  type = string
}

variable "region" {
  type = string
}

variable "type" {
  type = string
}

variable "flag" {
  type = bool
}

variable "port" {
  type = number
}

variable "tags" {
  type = list(string)
}

provider "scaleway" {
  zone   = var.zone
  region = var.region
}

resource "scaleway_instance_server" "web" {
  type = var.type
  zone = var.zone
  tags = var.tags
}

resource "scaleway_instance_security_group" "web" {
  inbound_default_policy = var.type
  tags                   = var.tags

  inbound_rule {
    action = "accept"
    port   = var.port
  }
}

resource "scaleway_k8s_cluster" "main" {
  name                        = "main"
  version                     = "1.35"
  cni                         = var.type
  delete_additional_resources = true
  tags                        = var.tags
}

resource "scaleway_k8s_pool" "default" {
  cluster_id = scaleway_k8s_cluster.main.id
  name       = "default"
  node_type  = var.type
  size       = 1
  tags       = var.tags
}

resource "scaleway_rdb_instance" "main" {
  name               = "main"
  node_type          = var.type
  engine             = var.type
  disable_backup     = var.flag
  encryption_at_rest = var.flag
  tags               = var.tags

  private_network {
    pn_id = "fr-par/11111111-1111-1111-1111-111111111111"
  }
}

resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = var.type
  user_name = "admin"
  password  = var.type
  tags      = var.tags

  private_network {
    id = "fr-par-1/11111111-1111-1111-1111-111111111111"
  }
}

resource "scaleway_lb" "main" {
  type                    = var.type
  ssl_compatibility_level = var.type
  tags                    = var.tags
}

resource "scaleway_object_bucket" "main" {
  name = "logs-${var.type}"
  tags = {
    env   = var.type
    owner = var.type
  }

  versioning {
    enabled = var.flag
  }
}

resource "scaleway_iam_policy" "main" {
  group_id = "11111111-1111-1111-1111-111111111111"
  tags     = var.tags

  rule {
    project_ids          = ["11111111-1111-1111-1111-111111111111"]
    permission_set_names = var.tags
  }
}
