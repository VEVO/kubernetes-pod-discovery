package server

import (
	"testing"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
)

func TestRun(t *testing.T) {
	endpointsCache := &cache.Endpoints{}
	go Run(8080, endpointsCache)
}
