package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	request, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "ok"
	if response.Body.String() != expected {
		t.Errorf("Unexpected response body: got %v want %v",
			response.Body.String(), expected)
	}
}
