#!/bin/bash

pushd $(dirname $0)

EXISTS=$(kind get clusters | grep devwiki-cluster)
if [ ! -z "$EXISTS" ]; then
    echo "Cluster already exists. Skipping cluster creation."
    exit 0
fi

# create kind cluster
kind create cluster -n devwiki-cluster --config ./manifest/kind-config.yaml --wait 60s

# install Gateway controller
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.7.1/standard-install.yaml
# sleep for a bit to let gateway api controller boot up. Ya I know this is bad.
sleep 10
helm install nginx-gateway oci://ghcr.io/nginxinc/charts/nginx-kubernetes-gateway --version 0.0.0-edge --create-namespace -n nginx-gateway --set nginxGateway.service.create=false

# switch nginx to host network routing (doesn't seem to be helm chart option for this atm)
kubectl patch deployment nginx-gateway-nginx-kubernetes-gateway -n nginx-gateway -p '{"spec":{ "template": {"spec": { "hostNetwork": true }}}}' --type=merge

# install Gateway
kubectl create namespace networking
kubectl apply -f ./manifest/cluster-gateway.yaml -n networking

# install metrics server
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
helm upgrade --install metrics-server metrics-server/metrics-server --namespace kube-system --set args[0]="--kubelet-insecure-tls=true"

echo Your Ready to Dev!

popd