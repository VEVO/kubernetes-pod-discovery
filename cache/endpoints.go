// Package cache provides access to a shared in-memory object cache for the endpoints
package cache

import (
	"sync"
	"time"

	"k8s.io/client-go/pkg/api/v1"
)

// Endpoints shared object that holds our cache of the endpoints. Requires the use of the RWMutex for locking.
type Endpoints struct {
	sync.RWMutex
	items       v1.Endpoints
	lastUpdated time.Time
}

// GetEndpoints gets the endpoints out of the cache
func (e *Endpoints) GetEndpoints() *v1.Endpoints {
	e.RLock()
	defer e.RUnlock()
	return &e.items
}

// GetLastUpdated gets the time of the last update to the cache
func (e *Endpoints) GetLastUpdated() *time.Time {
	e.RLock()
	defer e.RUnlock()
	return &e.lastUpdated
}

// SetEndpoints sets the endpoints on the cache
func (e *Endpoints) SetEndpoints(endpoints *v1.Endpoints) {
	e.Lock()
	defer e.Unlock()
	e.lastUpdated = time.Now()
	e.items = *endpoints
}

// Update watches a channel and updates the endpoints as items are pushed to the channel
func (e *Endpoints) Update(endpointsEvents chan *v1.Endpoints) {
	for endpoints := range endpointsEvents {
		e.SetEndpoints(endpoints)
	}
}
