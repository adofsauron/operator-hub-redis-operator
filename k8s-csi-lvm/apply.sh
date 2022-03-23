#!/bin/bash

# kubectl label nodes node-201 app=csi-lvmplugin

# kubectl taint nodes node-201 node-role.kubernetes.io/master-

kubectl describe nodes node-201

#
kubectl apply -f deploy/kubernetes/ 

kubectl delete -f deploy/kubernetes/ 

#
kubectl apply -f deploy/example/ 

kubectl delete -f deploy/example/ 

#
kubectl delete -f deploy/example/pod.yaml
kubectl apply -f deploy/example/pod.yaml
