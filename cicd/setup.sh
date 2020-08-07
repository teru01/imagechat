#!/bin/bash

set -exu
cd -- "${0%/*}"

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

while :; do
  ipaddr=$(kubectl get svc -n "$NAMESPACE" | awk '$2 == "LoadBalancer"{print $4}')
  if [ "$ipaddr" == "<pending>" ]; then
    echo -n "."
    sleep 1
  else
    break
  fi
done

current_password=$(kubectl get pods -n "$NAMESPACE" -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)

argocd login "$ipaddr" --username admin --password "$current_password"

new_password=$(awk -F"=" '$1 == "argocd-pass"{print $2}' ./.env)
argocd account update-password --current-password "$current_password" --new-password "$new_password"
