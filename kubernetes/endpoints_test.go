package kubernetes

import (
	"reflect"
	"testing"

	"github.com/VEVO/kubernetes-pod-discovery/fake"
)

func TestKubernetesClientConfig_GetEndpoints(t *testing.T) {
	endpointsController := NewEndpointsController(NewFakeClient())
	endpoints, err := endpointsController.GetEndpoints("service", "namespace")
	if err != nil {
		t.Errorf("Got error when fetching endpoints: %s", err)
	}
	if reflect.DeepEqual(endpoints, fake.Endpoints) {
		t.Errorf("Failed to get endpoints")
	}
}
