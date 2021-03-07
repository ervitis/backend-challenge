package endpoint

import (
	"context"
	"github.com/ervitis/backend-challenge/basket/domain"
	protopb "github.com/ervitis/backend-challenge/proto"
)

type (
	IRouter interface {
		GetRouter() protopb.BasketServer

		Create(context.Context, *protopb.NewOrderByUser) (*protopb.NewOrderCreated, error)
		Checkout(context.Context, *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error)
		AddToBasket(context.Context, *protopb.AddItemsToOrder) (*protopb.OperationConfirmed, error)
		GetTotalAmount(context.Context, *protopb.NewOrderCreated) (*protopb.Amount, error)
		RemoveAll(context.Context, *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error)
	}

	router struct {
		r protopb.BasketServer

		service domain.IBasket
	}
)

func (rt *router) GetRouter() protopb.BasketServer {
	return rt.r
}

func NewRouter(svc domain.IBasket) IRouter {
	return &router{service: svc}
}
