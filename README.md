# kubernetes-pod-discovery

This service can be used to cache the endpoints associated with a Kubernetes service. See https://kubernetes.io/docs/user-guide/services/.

Run this as a sidecar container and access it from your service via http://localhost:8080/endpoints or use it as an example to communicate
directly with the Kubernetes API from inside your process.

## Building

`go build .`

## Running

This service is designed to be run inside a Kubernetes Pod as a sidecar container. To run the example nginx service, apply the the following
configurations:

`kubectl apply -f examples/*.yaml`

This will create an nginx deployment and service which has type `LoadBalancer` so it will be accessible from the internet.
