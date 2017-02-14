// Package kubernetes provides access to the Kubernetes API
package kubernetes

import (
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
)

// EndpointsController provides access to the kubernetes endpoints functions
type EndpointsController struct {
	client Client
}

// NewEndpointsController stores our real or mocked kubernetes client object
func NewEndpointsController(client Client) *EndpointsController {
	return &EndpointsController{
		client: client,
	}
}

// GetEndpoints is a wrapper function around the kubernetes client. If there is any manipulation or filtering
// of the kubernetes types or data, it should be done here.
func (e EndpointsController) GetEndpoints(service string, namespace string) (*v1.Endpoints, error) {
	endpoints, err := e.client.GetEndpoints(service, namespace)
	if err != nil {
		err = fmt.Errorf("Could not get endpoints: %s", err)
		return &v1.Endpoints{}, err
	}
	return endpoints, nil
}
