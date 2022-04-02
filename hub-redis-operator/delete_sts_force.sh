#!/bin/bash

kubectl delete --force sts redis-cluster-leader

kubectl delete --force sts redis-cluster-follower

