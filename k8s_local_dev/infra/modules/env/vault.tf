resource "helm_release" "vault" {
  name = "vault"
  chart     = "hashicorp/vault"
  namespace = "default"

  provisioner "local-exec" {
    command = "sh infra/scripts/get-vault-ip.sh"
 
  }

  set {
    name  = "server.service.type"
    value = "NodePort"
  }

  set {
    name  = "server.service.nodePort"
    value = "31200"
  }

  set {
    name  = "server.dev.enabled"
    value = "true"
  }

  set {
    name  = "ui.enabled"
    value = "true"
  }
}