package backend

import (
	"errors"
	"sync"
)

type ServerPool struct {
	Backends []*Backend
	Mutex    sync.Mutex
}

func (serverPool *ServerPool) AddBackend(url string) error {
	serverPool.Mutex.Lock()
	defer serverPool.Mutex.Unlock()

	for _, backend := range serverPool.Backends {
		if backend.URL == url {
			return errors.New("backend already exists")
		}
	}

	serverPool.Backends = append(serverPool.Backends, &Backend{URL: url, Alive: true})

	return nil
}

func (serverPool *ServerPool) ListBackends() []*Backend {
	serverPool.Mutex.Lock()
	defer serverPool.Mutex.Unlock()
	return serverPool.Backends
}
