resource "helm_release" "postgres" {
  chart = "bitnami/postgresql"
  name = "postgres"

  
  values = [templatefile("${path.module}/values/postgres_helper.yaml.tpl", {
    consul_ip_val = "${var.consul_ip}"
    vault_ip_val = "${var.vault_ip}"
  })]

  set {
    name = "image.repository"
    value = "postgres"
  }

  set {
    name = "image.tag"
    value = "9.6"
  }

  set {
    name = "service.type"
    value = "NodePort"
  }

  set {
    name = "postgresqlUsername"
    value = "dbuser"
  }

  set {
    name = "postgresqlPassword"
    value = "p@$$w0rd"
  }  
}
