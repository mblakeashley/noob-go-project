
provider "helm" {
  kubernetes {
  }
}

resource "helm_release" "mydatabase" {
  name  = "mydatabase"
  chart = "bitnami/mariadb"

  set {
    name  = "mariadbUser"
    value = "foo"
  }

  set {
    name  = "mariadbPassword"
    value = "qux"
  }
}