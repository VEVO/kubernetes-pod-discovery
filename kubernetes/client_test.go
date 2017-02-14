package kubernetes

import (
	"github.com/VEVO/kubernetes-pod-discovery/fake"
	"k8s.io/client-go/pkg/api/v1"
)

type FakeKubernetesClientConfig struct{}

func NewFakeClient() Client {
	return &FakeKubernetesClientConfig{}
}

func (c FakeKubernetesClientConfig) GetEndpoints(service string, namespace string) (*v1.Endpoints, error) {
	return &fake.Endpoints, nil
}
