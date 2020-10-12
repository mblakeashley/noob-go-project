resource "helm_release" "consul" {
  name = "consul"
  chart     = "hashicorp/consul"
  namespace = "default"

  provisioner "local-exec" {
    command = "sh infra/scripts/get-consul-ip.sh"
  }

  set {
    name  = "server.replicas"
    value = 1
  }

  set {
    name  = "server.bootstrapExpect"
    value = 1
  }

  set {
    name  = "dns.enabled"
    value = true
  }

  set {
    name  = "ui.service.type"
    value = "NodePort"
  }

  set {
    name  = "server.connect"
    value = true
  }
}
