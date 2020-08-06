#!/bin/bash

set -exu

NAMESPACE=argocd

if bash -c "kubectl get namespace $NAMESPACE"; then
  :
else
  kubectl create namespace $NAMESPACE
fi

kubectl apply -n $NAMESPACE -f argocd-install.yaml

kubectl create secret generic repository-secret --from-file=.env \
  --dry-run -o yaml | kubectl apply -n $NAMESPACE -f -
kubectl apply -n $NAMESPACE -f argocd-setup.yaml
