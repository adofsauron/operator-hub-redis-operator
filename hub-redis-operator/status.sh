#!/bin/bash

echo kubectl get pods -o wide 
kubectl get pods -o wide | grep redis-cluster

echo -e "\n"

echo kubectl  exec redis-cluster-leader-0 -- redis-cli cluster nodes
kubectl  exec redis-cluster-leader-0 -- redis-cli cluster nodes

echo -e "\n"
