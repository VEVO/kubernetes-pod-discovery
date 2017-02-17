package main

import (
	"github.com/VEVO/kubernetes-pod-discovery/cache"
	"github.com/VEVO/kubernetes-pod-discovery/config"
	"github.com/VEVO/kubernetes-pod-discovery/kubernetes"
	"github.com/VEVO/kubernetes-pod-discovery/server"
	"k8s.io/client-go/pkg/api/v1"

	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Initialize the configuration
	conf := &config.Config{
		// ListenPort default
		ListenPort: 8080,
	}
	conf.Service = os.Getenv("KUBERNETES_POD_DISCOVERY_SERVICE_NAME")
	conf.Namespace = os.Getenv("KUBERNETES_POD_DISCOVERY_NAMESPACE")
	listenPort := os.Getenv("KUBERNETES_POD_DISCOVERY_LISTEN_PORT")
	if listenPort != "" {
		conf.ListenPort, _ = strconv.ParseInt(listenPort, 10, 64)
	}
	if err := conf.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Channel that holds our endpoint updates
	endpointEvents := make(chan *v1.Endpoints, 1)

	// Initialize the kubernetes client
	endpointController := kubernetes.NewEndpointsController(kubernetes.NewClient())

	// Watch for changes to the endpoints at the interval specified by the ticker
	ticker := time.NewTicker(time.Second * 5)
	go func(endpointController *kubernetes.EndpointsController, endpointEvents chan *v1.Endpoints, conf *config.Config) {
		fmt.Println(fmt.Sprintf("Watching endpoints for service %s, namespace %s", conf.Service, conf.Namespace))
		for range ticker.C {
			endpoints, err := endpointController.GetEndpoints(conf.Service, conf.Namespace)
			if err != nil {
				fmt.Println(fmt.Sprintf("Failed to get endpoints: %s", err))
				continue
			}
			endpointEvents <- endpoints
		}
	}(endpointController, endpointEvents, conf)

	// Feed endpoints updates to the cache
	endpointsCache := &cache.Endpoints{}
	go endpointsCache.Update(endpointEvents)

	// Serve up an http endpoint to access the endpoint cache
	server.Run(conf.ListenPort, endpointsCache)
}
