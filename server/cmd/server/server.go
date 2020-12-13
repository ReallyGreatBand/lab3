package main

import (
	"context"
	"fmt"
	"lab3/server/plants"
	"net/http"
)

type HttpPortNumber int

// PlantApiServer configures necessary handlers and starts listening on a configured port.
type PlantApiServer struct {
	Port HttpPortNumber

	PlantsHandler plants.HttpHandlerFunc

	server *http.Server
}

func (s *PlantApiServer) Start() error {
	if s.PlantsHandler == nil {
		return fmt.Errorf("plant's HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/plants", s.PlantsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}
	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *PlantApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}

	return s.server.Shutdown(context.Background())
}
