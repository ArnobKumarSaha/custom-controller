#!/bin/bash

kind create cluster --name first

go build -o app .

kubectl apply -f=manifests/arnob.com_messis.yaml

./app -kubeconfig=$HOME/.kube/config
