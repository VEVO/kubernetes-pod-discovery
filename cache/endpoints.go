package cache

import (
	"sync"
	"time"

	"k8s.io/client-go/pkg/api/v1"
)

// Shared object that holds our cache of the endpoints. Requires the use of the RWMutex for locking.
type EndpointsCache struct {
	lock        sync.RWMutex
	items       v1.Endpoints
	lastUpdated time.Time
}

// Get the endpoints out of the cache
func (e *EndpointsCache) GetEndpoints() *v1.Endpoints {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return &e.items
}

// Get the time of the last update to the cache
func (e *EndpointsCache) GetLastUpdated() *time.Time {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return &e.lastUpdated
}

// Set the endpoints on the cache
func (e *EndpointsCache) SetEndpoints(endpoints *v1.Endpoints) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.lastUpdated = time.Now()
	e.items = *endpoints
}

// Watch a channel and update the endpoints as items are pushed to the channel
func (e *EndpointsCache) Update(endpointsEvents chan *v1.Endpoints) {
	for endpoints := range endpointsEvents {
		e.SetEndpoints(endpoints)
	}
}
