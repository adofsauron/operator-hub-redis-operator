#!/bin/bash

export KUBERNETES_SERVICE_HOST="127.0.0.1"
export KUBERNETES_SERVICE_PORT="6443"

./ctr -development=true -debug=true