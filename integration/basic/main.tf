provider "scaleway" {
  zone   = "fr-par-1"
  region = "fr-par-1"
}

resource "scaleway_instance_server" "web" {
  type  = "DEV1-XXL"
  image = "ubuntu_noble"
  zone  = "fr-par-4"
}

resource "scaleway_k8s_cluster" "main" {
  name                        = "main"
  version                     = "1.35"
  cni                         = "flannel"
  delete_additional_resources = true
}

resource "scaleway_k8s_pool" "default" {
  cluster_id = scaleway_k8s_cluster.main.id
  name       = "default"
  node_type  = "t2.micro"
  size       = 3
}

resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "db-dev-xxs"
  engine    = "postgres-16"
}

resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = "RED1-XXL"
  user_name = "admin"
  password  = var.redis_password
}

resource "scaleway_lb" "main" {
  name = "main"
  type = "LB-XXL"
}

resource "scaleway_object_bucket" "logs" {
  name = "My_Logs"
}

# Valid configuration: no issue expected below.
resource "scaleway_instance_server" "ok" {
  type  = "pro2-xxs"
  image = "ubuntu_noble"
  zone  = "nl-ams-1"
}

variable "redis_password" {
  type      = string
  sensitive = true
}
