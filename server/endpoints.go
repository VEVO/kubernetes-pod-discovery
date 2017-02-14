package server

import (
	"encoding/json"
	"io"
	"net/http"

	"fmt"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
)

// Our interface that specifies the endpoint routes available for our http server
type Endpoints interface {
	Root(http.ResponseWriter, *http.Request)
	LastUpdated(http.ResponseWriter, *http.Request)
}

// Object that we use to store access to the cache through our endpoints routes
type EndpointsServer struct {
	cache *cache.EndpointsCache
}

// Create a new endpoints server and point to the specified cache
func NewEndpointsServer(endpointsCache *cache.EndpointsCache) Endpoints {
	return &EndpointsServer{
		cache: endpointsCache,
	}
}

// Serve our root endpoints route
func (e *EndpointsServer) Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	endpoints, err := json.Marshal(*e.cache.GetEndpoints())
	if err != nil {
		io.WriteString(w, fmt.Sprintf("%s", err))
	}
	io.WriteString(w, string(endpoints))
}

// Serve our last_updated endpoints route
func (e *EndpointsServer) LastUpdated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lastUpdated := *e.cache.GetLastUpdated()
	response := fmt.Sprintf("{\"lastUpdated\": \"%s\"}", lastUpdated.UTC().Format("2006-01-02T15:04:05-0700"))
	io.WriteString(w, string(response))
}
