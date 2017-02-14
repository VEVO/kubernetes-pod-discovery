// Package kubernetes provides access to the Kubernetes API
package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	meta_v1 "k8s.io/client-go/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Client interface that allows us to mock calls to the Kubernetes API
type Client interface {
	GetEndpoints(service string, namespace string) (*v1.Endpoints, error)
}

// ClientConfig holds the kubernetes session for when communicating with the kubernetes API
type ClientConfig struct {
	clientset *kubernetes.Clientset
}

// NewClient initiates connection to the kubernetes API
// See https://github.com/kubernetes/client-go/blob/master/examples/in-cluster/main.go
func NewClient() Client {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &ClientConfig{clientset: clientset}
}

// GetEndpoints gets the endpoints from the kubernetes API
// See https://github.com/kubernetes/client-go
func (c ClientConfig) GetEndpoints(service string, namespace string) (*v1.Endpoints, error) {
	endpointsGetter := c.clientset.Core().Endpoints(namespace)
	return endpointsGetter.Get(service, meta_v1.GetOptions{})
}
