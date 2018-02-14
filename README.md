# kubernetes-pod-discovery

[![Build Status](https://travis-ci.org/VEVO/kubernetes-pod-discovery.svg?branch=master)](https://travis-ci.org/VEVO/kubernetes-pod-discovery)
[![Go Report Card](https://goreportcard.com/badge/github.com/VEVO/kubernetes-pod-discovery)](https://goreportcard.com/report/github.com/VEVO/kubernetes-pod-discovery)

This service can be used to cache the endpoints associated with a Kubernetes service. See https://kubernetes.io/docs/user-guide/services/.

Run this as a sidecar container and access it from your service via http://localhost:8080/endpoints or use it as an example to communicate directly with the Kubernetes API from inside your process.

## How It Works

This service runs inside the Kubernetes cluster (typically as a sidecar container) and uses the built-in service account token to authenticate with the Kubernetes API (https://kubernetes.io/docs/user-guide/accessing-the-cluster/#accessing-the-api-from-a-pod). It caches the endpoints for a given service and namespace, which are specified using the environment variables `KUBERNETES_POD_DISCOVERY_SERVICE_NAME` and `KUBERNETES_POD_DISCOVERY_NAMESPACE` and presents them to the following endpoint that will be accessible from any containers in the pod:

`http://localhost:8080/endpoints`

There is also a health endpoint that can be accessed from the following:

`http://localhost:8080/health`

Note that the listen port can be adjusted via the `KUBERNETES_POD_DISCOVERY_LISTEN_PORT` environment variable.

## Format

The endpoints format is the following:

```
{
  metadata: <kubernetes metadata>...
  },
  subsets: [
  {
    addresses: [
    {
      ip: "172.20.17.5",
      nodeName: "ip-10-20-9-161.ec2.internal",
      targetRef: { <more metadata>
      }
    },
    ],
    ports: [
    {
      name: <service name>,
      port: 80,
      protocol: "TCP"
    },
    ]
  }
  ]
}
```

So with the above configuration, the addresses for the available pods in the service pool can be found from `subsets.0.addresses.0.ip` where 0 is the index of the list. The ports can be accessed the same way.

## Building

`go build .`

## Running

This service is designed to be run inside a Kubernetes Pod as a sidecar container. To run the example nginx service, apply the the following configurations:

`kubectl apply -R -f examples/`

This will create an nginx deployment and service which has type `LoadBalancer` so it will be accessible from the internet.

## Pre-built Docker Image

You can use the pre-built docker image at https://hub.docker.com/r/vevo/kubernetes-pod-discovery. For example, your sidecar container would use image: `vevo/kubernetes-pod-discovery:0.0.3`
