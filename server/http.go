package server

import (
	"fmt"
	"net/http"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
)

// Runs our http server using specified port and cache
func Run(port int64, endpointsCache *cache.EndpointsCache) {
	endpointsServer := NewEndpointsServer(endpointsCache)

	http.HandleFunc("/endpoints/", endpointsServer.Root)
	http.HandleFunc("/endpoints/last_updated/", endpointsServer.LastUpdated)
	http.HandleFunc("/health/", Health)

	fmt.Println(fmt.Sprintf("Server started on port %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
