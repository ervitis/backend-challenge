package main

import (
	"github.com/ervitis/backend-challenge/basket/domain"
	"github.com/ervitis/backend-challenge/basket/endpoint"
	"github.com/ervitis/backend-challenge/basket/repository"
	"github.com/ervitis/backend-challenge/basket/server"
	"github.com/ervitis/logme"
	"github.com/ervitis/logme/config_loaders"
	"os"
)

func init() {
	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	_ = os.Setenv("LOG_FIELDS", "[component=server,service=basket]")
}

func main() {
	loggerConfig, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	logger := logme.NewLogme(loggerConfig)

	router := endpoint.NewRouter(domain.NewBasketService(repository.NewRepository()))
	srv := server.CreateServer(server.WithRouter(router), server.WithPort("8181"), server.WithLogger(logger))
	srv.Listen()
}
