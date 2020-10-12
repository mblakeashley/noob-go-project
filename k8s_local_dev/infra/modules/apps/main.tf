data "local_file" "consul-ip" {
  filename = "${path.module}/outputs/consul_ips"
}

data "local_file" "vault-ip" {
  filename = "${path.module}/outputs/vault_ips"
}

output "consul_ips" {
  value = "${data.local_file.consul-ip.content}"
}

output "vault_ips" {
  value = "${data.local_file.vault-ip.content}"
} 