// Package server serves up the endpoints cache via http
package server

import (
	"fmt"
	"net/http"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
)

// Run runs our http server using specified port and cache
func Run(port int64, endpointsCache *cache.Endpoints) {
	endpointsServer := NewEndpointsServer(endpointsCache)

	http.HandleFunc("/endpoints/", endpointsServer.Root)
	http.HandleFunc("/endpoints/last_updated/", endpointsServer.LastUpdated)
	http.HandleFunc("/health/", Health)

	fmt.Println(fmt.Sprintf("Server started on port %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
