package backend

import "sync"

type Backend struct {
	URL   string
	Alive bool
	Mutex sync.Mutex
}

func (backend *Backend) MarkAlive(alive bool) {
	backend.Mutex.Lock()
	defer backend.Mutex.Unlock()
	backend.Alive = alive
}
