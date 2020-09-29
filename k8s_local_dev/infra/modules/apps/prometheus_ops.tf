module "release-prometheus-operator" {
  source  = "OpenQAI/release-prometheus-operator/helm"
  version = "0.0.6"

  helm_chart_version    = "8.15.11"
  helm_chart_namespace  = "monitoring"
  skip_crds             = false
  prometheus_service_type = "NodePort"
  grafana_image_tag      = "7.0.3"
  grafana_adminPassword  = "pa$$w0rd"
  grafana_service_type  = "NodePort"
}
