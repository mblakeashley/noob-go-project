provider "helm" {
  kubernetes {
  }
}

module "apps" {
  source = "./modules/apps"
  depends_on = [module.env]
  consul_ip = module.apps.consul_ips
  vault_ip = module.apps.vault_ips
}

module "env" {
  source = "./modules/env"
}

output "consul_ips" {
  value = module.apps.consul_ips
}

output "vault_ips" {
  value = module.apps.vault_ips
}