resource "helm_release" "consul" {
  name = "consul"
  chart     = "hashicorp/consul"
  namespace = "default"

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
