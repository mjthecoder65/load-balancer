package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mjthecoder65/load-balancer/backend"
	"github.com/mjthecoder65/load-balancer/balancer"
	"github.com/mjthecoder65/load-balancer/proxy"
)

func main() {
	backends := []*backend.Backend{
		{URL: "http://localhost:8081", Alive: true},
		{URL: "http://localhost:8082", Alive: true},
	}

	roundRobin := &balancer.RoundRobin{Backends: backends}

	go func() {
		for {
			for _, backend := range backends {
				backend.CheckHealth()
			}
			time.Sleep(10 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backend := roundRobin.NextBackend()

		if backend == nil {
			http.Error(w, "no backend available", http.StatusServiceUnavailable)
			return
		}
		proxy.ForwardRequest(backend, w, r)
	})

	fmt.Println("Load Balancer Started Running on port :8080")
	http.ListenAndServe(":8080", nil)
}
