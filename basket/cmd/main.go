package main

import (
	"github.com/ervitis/backend-challenge/basket/domain"
	"github.com/ervitis/backend-challenge/basket/endpoint"
	"github.com/ervitis/backend-challenge/basket/repository"
	"github.com/ervitis/backend-challenge/basket/server"
)

func main() {
	router := endpoint.NewRouter(domain.NewBasketService(repository.NewRepository()))
	srv := server.CreateServer(server.WithRouter(router))
	srv.Listen()
}
