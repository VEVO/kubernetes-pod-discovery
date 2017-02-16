package config

import (
	"fmt"
	"testing"
)

var (
	service   = "service"
	namespace = "namespace"
)

func TestConfig_Validate(t *testing.T) {
	config := &Config{}
	if err := config.Validate(); err != nil {
		if fmt.Sprintf("%s", err) != "The service is required" {
			t.Errorf("Expected error but got %s", err)
		}
	} else {
		t.Errorf("Expected error but didn't get one")
	}
}

func TestConfig_Validate2(t *testing.T) {
	config := &Config{
		Service: service,
	}
	if err := config.Validate(); err != nil {
		if fmt.Sprintf("%s", err) != "The namespace is required" {
			t.Errorf("Expected error but got %s", err)
		}
	} else {
		t.Errorf("Expected error but didn't get one")
	}
}

func TestConfig_Validate3(t *testing.T) {
	config := &Config{
		Service:   service,
		Namespace: namespace,
	}
	if err := config.Validate(); err != nil {
		t.Errorf("Did not expect error but got %s", err)
	}
}
