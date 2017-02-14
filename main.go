package main

import (
	"github.com/VEVO/kubernetes-pod-discovery/cache"
	"github.com/VEVO/kubernetes-pod-discovery/config"
	"github.com/VEVO/kubernetes-pod-discovery/kubernetes"
	"github.com/VEVO/kubernetes-pod-discovery/server"
	"k8s.io/client-go/pkg/api/v1"

	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	conf := &config.Config{}
	flag.StringVar(&conf.Service, "service", "", "name of the service")
	flag.StringVar(&conf.Namespace, "namespace", "", "name of the namespace where the service resides")
	flag.Int64Var(&conf.ListenPort, "listen-port", 8080, "port on which to bind the service")
	flag.Parse()

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
				time.Sleep(time.Duration(1 * time.Second))
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
