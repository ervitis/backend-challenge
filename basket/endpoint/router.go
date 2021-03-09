package endpoint

import (
	"context"
	"github.com/ervitis/backend-challenge/basket/domain"
	protopb "github.com/ervitis/backend-challenge/proto"
)

type (
	IRouter interface {
		GetRouter() *router
		Create(context.Context, *protopb.NewOrderByUser) (*protopb.NewOrderCreated, error)
		Checkout(context.Context, *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error)
		AddToBasket(context.Context, *protopb.AddItemsToOrder) (*protopb.OperationConfirmed, error)
		GetTotalAmount(context.Context, *protopb.NewOrderCreated) (*protopb.Amount, error)
		RemoveAll(context.Context, *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error)
	}

	router struct {
		protopb.UnimplementedBasketServer
		service domain.IBasket
	}
)

func (rt *router) GetRouter() *router {
	return rt
}

func NewRouter(svc domain.IBasket) *router {
	return &router{service: svc}
}
