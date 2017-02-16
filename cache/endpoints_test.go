package cache

import (
	"reflect"
	"testing"
	"time"

	"github.com/VEVO/kubernetes-pod-discovery/fake"
	"k8s.io/client-go/pkg/api/v1"
)

func TestEndpoints_GetEndpoints(t *testing.T) {
	endpointsCache := Endpoints{}

	if len(endpointsCache.GetEndpoints().Subsets) != 0 {
		t.Error("Failed to get empty endpoints")
	}
}

func TestEndpoints_SetEndpoints(t *testing.T) {
	endpointsCache := Endpoints{}
	fakeEndpoints := &fake.Endpoints
	endpointsCache.SetEndpoints(fakeEndpoints)

	if !reflect.DeepEqual(endpointsCache.GetEndpoints(), fakeEndpoints) {
		t.Error("Failed to set endpoints")
	}
}

func TestEndpoints_Update(t *testing.T) {
	endpointsCache := Endpoints{}
	fakeEndpoints := &fake.Endpoints
	endpointEvents := make(chan *v1.Endpoints)

	go endpointsCache.Update(endpointEvents)

	endpointEvents <- fakeEndpoints
	if !reflect.DeepEqual(endpointsCache.GetEndpoints(), fakeEndpoints) {
		t.Error("Failed to set endpoints")
	}
}

func TestEndpoints_GetLastUpdated(t *testing.T) {
	endpointsCache := Endpoints{}
	fakeEndpoints := &fake.Endpoints
	timeMarker := time.Now()
	endpointsCache.SetEndpoints(fakeEndpoints)
	if endpointsCache.GetLastUpdated().Before(timeMarker) {
		t.Error("Failed to get last updated time")
	}
}
