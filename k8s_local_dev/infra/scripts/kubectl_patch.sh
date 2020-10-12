# Set of patches to set NodePort values
kubectl patch service prometheus-operator-grafana --namespace=monitoring --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":31448}]'
kubectl patch service prometheus-operator-prometheus --namespace=monitoring --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":30090}]'
kubectl patch service consul-consul-ui --namespace=default --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":31349}]'
kubectl patch service postgres-postgresql --namespace=default --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":31902}]'
kubectl patch service vault --namespace=default --type='json' --patch='[{"op": "replace", "path": "/spec/ports/0/nodePort", "value":31200}]'