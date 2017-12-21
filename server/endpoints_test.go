package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
	"github.com/VEVO/kubernetes-pod-discovery/fake"
)

func TestEndpointsServer_Root(t *testing.T) {
	endpoints := &fake.Endpoints
	endpointsCache := &cache.Endpoints{}
	endpointsCache.SetEndpoints(endpoints)
	endpointsServer := NewEndpointsServer(endpointsCache)

	request, err := http.NewRequest("GET", "/endpoints", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(endpointsServer.Root)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if expected, err := json.Marshal(endpoints); response.Body.String() != string(expected) {
		t.Errorf("Unexpected response body: got %v want %v",
			response.Body.String(), expected)
	} else {
		if err != nil {
			t.Errorf("Unexpected json Marshalling error: %s", err)
		}
	}
}

func TestEndpointsServer_LastUpdated(t *testing.T) {
	endpoints := &fake.Endpoints
	endpointsCache := &cache.Endpoints{}
	endpointsCache.SetEndpoints(endpoints)
	endpointsServer := NewEndpointsServer(endpointsCache)

	request, err := http.NewRequest("GET", "/endpoints/last_updated", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(endpointsServer.LastUpdated)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := fmt.Sprintf("{\"lastUpdated\": \"%s\"}", endpointsCache.GetLastUpdated().UTC().Format("2006-01-02T15:04:05-0700"))
	if response.Body.String() != string(expected) {
		t.Errorf("Unexpected response body: got %v want %v",
			response.Body.String(), expected)
	}
}
