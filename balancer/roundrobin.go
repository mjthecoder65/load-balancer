package balancer

import (
	"sync"

	"github.com/mjthecoder65/load-balancer/backend"
)

type RoundRobin struct {
	Backends []*backend.Backend
	Current  int
	Mutex    sync.Mutex
}

func (roundRobin *RoundRobin) NextBackend() *backend.Backend {
	roundRobin.Mutex.Lock()
	defer roundRobin.Mutex.Unlock()

	serversCounts := len(roundRobin.Backends)

	for i := 0; i < serversCounts; i++ {
		roundRobin.Current = (roundRobin.Current + 1) % serversCounts
		if roundRobin.Backends[roundRobin.Current].Alive {
			return roundRobin.Backends[roundRobin.Current]
		}
	}
	return nil
}
