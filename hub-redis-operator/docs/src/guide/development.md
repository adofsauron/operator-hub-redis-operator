# Development Guide

## Pre-requisites

**Access to Kubernetes cluster**

First of all, you will need access to a Kubernetes cluster. The easiest way to start is minikube.

- [Virtualbox](https://www.virtualbox.org/wiki/Downloads) - hypervisor to run a Kubernetes cluster
- [Minikube](https://kubernetes.io/docs/setup/minikube/) - for Kubernetes cluster creation on local machine
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - to interact with Kubernetes cluster

**Tools to build an Operator**

Apart from kubernetes cluster, there are some tools which are needed to build and test the redis operator.

- [Git](https://git-scm.com/downloads)
- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)
- [Operator SDK](https://github.com/operator-framework/operator-sdk/blob/v0.8.1/doc/user/install-operator-sdk.md)
- [Make](https://www.gnu.org/software/make/manual/make.html)

## Build

**Build Locally**

To achieve this, execute this command:-

```shell
$ make manager
...
/go/src/redis-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
```

**Build Image**

Redis operator gets packaged as a container image for running on the Kubernetes cluster. These instructions will guide you to build an image.

```shell
$ make docker-build
```

## Testing

If you want to play it on Kubernetes. You can use a minikube.

```shell
$ minikube start --vm-driver virtualbox
...
๐  minikube v1.0.1 on linux (amd64)
๐คน  Downloading Kubernetes v1.14.1 images in the background ...
๐ฅ  Creating kvm2 VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
๐ถ  "minikube" IP address is 192.168.39.240
๐ณ  Configuring Docker as the container runtime ...
๐ณ  Version of container runtime is 18.06.3-ce
โ  Waiting for image downloads to complete ...
โจ  Preparing Kubernetes environment ...
๐  Pulling images required by Kubernetes v1.14.1 ...
๐  Launching Kubernetes v1.14.1 using kubeadm ... 
โ  Waiting for pods: apiserver proxy etcd scheduler controller dns
๐  Configuring cluster permissions ...
๐ค  Verifying component health .....
๐  kubectl is now configured to use "minikube"
๐  Done! Thank you for using minikube!
```

**Running Test Cases**

```shell
$ make test
```