package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	meta_v1 "k8s.io/client-go/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Kubernetes client interface that allows us to mock calls to the kubernetes API
type KubernetesClient interface {
	GetEndpoints(service string, namespace string) (*v1.Endpoints, error)
}

// Holds the kubernetes session for when communicating with the kubernetes API
type KubernetesClientConfig struct {
	clientset *kubernetes.Clientset
}

// Initiates connection to the kubernetes API
// See https://github.com/kubernetes/client-go/blob/master/examples/in-cluster/main.go
func NewClient() KubernetesClient {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &KubernetesClientConfig{clientset: clientset}
}

// Get the endpoints from the kubernetes API
// See https://github.com/kubernetes/client-go
func (c KubernetesClientConfig) GetEndpoints(service string, namespace string) (*v1.Endpoints, error) {
	endpointsGetter := c.clientset.Core().Endpoints(namespace)
	return endpointsGetter.Get(service, meta_v1.GetOptions{})
}
