# Set of patches to set NodePort values

kubectl patch service prometheus-operator-grafana --namespace=monitoring --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":31448}]'