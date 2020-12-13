//+build wireinject

package main

import (
	"lab3/server/plants"

	"github.com/google/wire"
)

// ComposeApiServer will create an instance of PlantApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*PlantApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from plants package.
		plants.Providers,
		// Provide PlantApiServer instantiating the structure and injecting plants handler and port number.
		wire.Struct(new(PlantApiServer), "Port", "PlantsHandler"),
	)
	return nil, nil
}
