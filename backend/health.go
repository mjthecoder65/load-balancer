package backend

import (
	"net/http"
	"time"
)

func (backend *Backend) CheckHealth() {
	timeout := 2 * time.Second
	client := http.Client{Timeout: timeout}

	res, err := client.Get(backend.URL + "/health")
	backend.MarkAlive(err == nil && res.StatusCode == http.StatusOK)
}
