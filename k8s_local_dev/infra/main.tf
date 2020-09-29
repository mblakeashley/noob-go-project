provider "helm" {
  kubernetes {
  }
}

module "apps" {
  source = "./modules/apps"
}
