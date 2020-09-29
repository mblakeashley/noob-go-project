# K8s_Local_Dev

K8s_Local_Dev is a fully automated Kubernetes deployment written in Go. The support toolchain consists of [Golang](https://golang.org/), [Terraform](https://www.terraform.io/), [Docker](https://www.docker.com/) and [Kind](https://kind.sigs.k8s.io/) (local K8s distro).


**Dependencies**

 1. Linx Distro/MacOS
 2. Docker
 3. Golang
 4. Terraform
 5. Helm
 6. Kind
 7. Make

**Example Output**
```
~$ ./main             
go run main.go
Kind cluster exists, will delete cluster...
Deploying New Cluster... 
 
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.19.1) 🖼 
 ✓ Preparing nodes 📦 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
 ✓ Joining worker nodes 🚜 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Thanks for using kind! 😊

+ -------- Starting Infra Setup -------- +
- Running Make Build...
+ Make Build Completed!
- Running Terraform Init... 
+ Init Completed!
- Running Terraform Apply... (Grab a cup of coffee, this will take a few minutes to complete.) 
+ Apply Completed!
- Applying Kube Patch Configs... 
+ -------- Infra Setup Complete -------- +

Run: 'kubectl get pods --all-namespaces' to view pods that are running in the cluster

~$ kubectl get pods --all-namespaces
NAMESPACE            NAME                                                     READY   STATUS    RESTARTS   AGE
kube-system          coredns-f9fd979d6-2lzbg                                  1/1     Running   0          10m
kube-system          coredns-f9fd979d6-q5ffp                                  1/1     Running   0          10m
kube-system          etcd-kind-control-plane                                  1/1     Running   0          10m
kube-system          kindnet-lqm85                                            1/1     Running   0          10m
kube-system          kindnet-p4hlc                                            1/1     Running   0          10m
kube-system          kube-apiserver-kind-control-plane                        1/1     Running   0          10m
kube-system          kube-controller-manager-kind-control-plane               1/1     Running   0          10m
kube-system          kube-proxy-5pmnn                                         1/1     Running   0          10m
kube-system          kube-proxy-gkskm                                         1/1     Running   0          10m
kube-system          kube-scheduler-kind-control-plane                        1/1     Running   0          10m
local-path-storage   local-path-provisioner-78776bfc44-mpjz6                  1/1     Running   0          10m
monitoring           alertmanager-prometheus-operator-alertmanager-0          2/2     Running   0          9m21s
monitoring           prometheus-operator-grafana-69876d8544-xxbw5             2/2     Running   0          9m40s
monitoring           prometheus-operator-kube-state-metrics-bd8f49464-db8zl   1/1     Running   0          9m40s
monitoring           prometheus-operator-operator-657c8fc5d6-vttmw            2/2     Running   0          9m40s
monitoring           prometheus-operator-prometheus-node-exporter-n6dfk       1/1     Running   0          9m40s
monitoring           prometheus-operator-prometheus-node-exporter-zbwgl       1/1     Running   0          9m40s
monitoring           prometheus-prometheus-operator-prometheus-0              3/3     Running   1          9m11s
```


**Launch Instructions**

 1. Clone repo and cd to k8s_local_dev directory
 ```git clone git@github.com:mblakeashley/noob-go-project.git && cd noob-go-project/k8s_local_dev ```
 2. Install listed dependencies
 3.  Run the Go binary `./main` from within the k8s_local_dev directory
