package main

import (
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"github.com/ervitis/backend-challenge/clientrest/endpoint"
	"github.com/ervitis/backend-challenge/clientrest/infra"
	"github.com/ervitis/backend-challenge/clientrest/rpccaller"
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
	rpcClient := rpccaller.New(infra.App.BasketAddr, infra.Logger)
	if err := rpcClient.Open(); err != nil {
		panic(err)
	}
	defer rpcClient.Close()

	router := endpoint.NewRouter(domain.NewBasketService(rpcClient))
	router.LoadApi()
	srv := server.CreateServer(server.WithRouter(router), server.WithLogger(infra.Logger))

	srv.Listen()
}
