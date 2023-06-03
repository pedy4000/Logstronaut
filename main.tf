# main.tf

provider "digitalocean" {
  token = "YOUR_DIGITALOCEAN_API_TOKEN"
}

resource "digitalocean_kubernetes_cluster" "logstronaut_cluster" {
  name    = "logstronaut_cluster"
  region  = "ams3"
  version = "latest"

  node_pool {
    name       = "worker-pool-1"
    size       = "s-1vcpu-2gb"
    node_count = 3
  }

  node_pool {
    name       = "worker-pool-2"
    size       = "s-4vcpu-8gb"
    node_count = 1
  }
}
