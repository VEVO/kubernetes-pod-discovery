package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
	"github.com/VEVO/kubernetes-pod-discovery/fake"
)

func TestEndpointsServer_Root(t *testing.T) {
	endpoints := &fake.Endpoints
	endpointsCache := &cache.EndpointsCache{}
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

	expected, err := json.Marshal(endpoints)
	if response.Body.String() != string(expected) {
		t.Errorf("Unexpected response body: got %v want %v",
			response.Body.String(), expected)
	}
}
