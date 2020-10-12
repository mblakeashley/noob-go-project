#!/bin/bash

vault_ips="infra/modules/apps/outputs/vault_ips"
if [ -f $vault_ips ] ; then
    rm $vault_ips > /dev/null 2>&1
fi

while ! cat $vault_ips > /dev/null 2>&1
do
kubectl get pods -l app.kubernetes.io/name=vault -o go-template='{{range .items}}{{.status.podIP}}{{"\n"}}{{end}}'  > ${vault_ips}
    sleep 1
done