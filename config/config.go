// Package config stores our config that is used by the service
package config

import "fmt"

// Config holds our configuration for the service
type Config struct {
	Service    string
	Namespace  string
	ListenPort int64
}

// Validate validates that the required flags are passed to the service
func (c *Config) Validate() error {
	if c.Service == "" {
		return fmt.Errorf("The service is required")
	}
	if c.Namespace == "" {
		return fmt.Errorf("The namespace is required")
	}
	return nil
}
