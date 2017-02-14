package fake

import "k8s.io/client-go/pkg/api/v1"

var NodeName = "node"
var Endpoints = v1.Endpoints{
	Subsets: []v1.EndpointSubset{
		v1.EndpointSubset{
			Addresses: []v1.EndpointAddress{
				v1.EndpointAddress{
					IP:       "127.0.0.1",
					Hostname: "localhost",
					NodeName: &NodeName,
				},
			},
			Ports: []v1.EndpointPort{
				v1.EndpointPort{
					Port: 8080,
				},
			},
		},
	},
}
