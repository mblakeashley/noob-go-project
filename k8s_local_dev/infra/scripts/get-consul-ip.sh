#!\bin\bash

consul_ips="infra/modules/apps/outputs/consul_ips"
if [ -f $consul_ips ] ; then
    rm $consul_ips > /dev/null 2>&1
fi

while ! cat $consul_ips > /dev/null 2>&1
do
kubectl get pods -l statefulset.kubernetes.io/pod-name=consul-consul-server-0 -o go-template='{{range .items}}{{.status.podIP}}{{"\n"}}{{end}}'  > $consul_ips
    sleep 1
done