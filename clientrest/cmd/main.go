package main

import (
	"github.com/ervitis/backend-challenge/clientrest/server"
	"github.com/ervitis/logme"
	"github.com/ervitis/logme/config_loaders"
	"net/http"
)

func main() {
	loggerConfig, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	logger := logme.NewLogme(loggerConfig)
	srv := server.CreateServer(server.WithRouter(http.NewServeMux()), server.WithLogger(logger))

	srv.Listen()
}
