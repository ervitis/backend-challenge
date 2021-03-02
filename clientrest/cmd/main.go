package main

import (
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"github.com/ervitis/backend-challenge/clientrest/endpoint"
	"github.com/ervitis/backend-challenge/clientrest/infra"
	"github.com/ervitis/backend-challenge/clientrest/server"
	"github.com/ervitis/logme"
	"github.com/ervitis/logme/config_loaders"
)

func init() {
	infra.LoadConfig()
}

func main() {
	loggerConfig, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	infra.Logger = logme.NewLogme(loggerConfig)
	router := endpoint.NewRouter(domain.NewBasketService(infra.Logger))
	router.LoadApi()
	srv := server.CreateServer(server.WithRouter(router), server.WithLogger(infra.Logger))

	srv.Listen()
}
