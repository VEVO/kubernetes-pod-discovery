// Package fake provides some mock objects
package fake

import "k8s.io/client-go/pkg/api/v1"

// NodeName provides a mock Kubernetes node name
var NodeName = "node"

// Endpoints provides mock Kubernetes endpoints
var Endpoints = v1.Endpoints{
	Subsets: []v1.EndpointSubset{
		{
			Addresses: []v1.EndpointAddress{
				{
					IP:       "127.0.0.1",
					Hostname: "localhost",
					NodeName: &NodeName,
				},
			},
			Ports: []v1.EndpointPort{
				{
					Port: 8080,
				},
			},
		},
	},
}
