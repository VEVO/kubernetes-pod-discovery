package config

import "fmt"

// Hold our configuration for the service
type Config struct {
	Service    string
	Namespace  string
	ListenPort int64
}

// Validate that the required flags are passed to the service
func (c *Config) Validate() error {
	if c.Service == "" {
		return fmt.Errorf("The service flag is required")
	}
	if c.Namespace == "" {
		return fmt.Errorf("The namespace flag is required")
	}
	return nil
}
