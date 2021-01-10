package main

import (
	"github.com/ervitis/backend-challenge/clientrest/endpoint"
	"github.com/ervitis/backend-challenge/clientrest/server"
	"github.com/ervitis/logme"
	"github.com/ervitis/logme/config_loaders"
)

func main() {
	loggerConfig, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	logger := logme.NewLogme(loggerConfig)
	router := endpoint.NewRouter()
	srv := server.CreateServer(server.WithRouter(router), server.WithLogger(logger))

	srv.Listen()
}
