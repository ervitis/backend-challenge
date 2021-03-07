package endpoint

import (
	"context"
	protopb "github.com/ervitis/backend-challenge/proto"
)

func (rt *router) Create(ctx context.Context, orderByUser *protopb.NewOrderByUser) (*protopb.NewOrderCreated, error) {
	orderID, err := rt.service.CreateOrder(ctx, orderByUser.UserID)
	return &protopb.NewOrderCreated{OrderID: orderID}, err
}

func (rt *router) Checkout(ctx context.Context, order *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error) {
	if err := rt.service.Checkout(ctx, order.OrderID); err != nil {
		return &protopb.OperationConfirmed{Ok: false}, err
	}
	return &protopb.OperationConfirmed{Ok: true}, nil
}

func (rt *router) AddToBasket(ctx context.Context, items *protopb.AddItemsToOrder) (*protopb.OperationConfirmed, error) {
	if err := rt.service.AddItemToBasket(ctx, items.OrderID, items.Items); err != nil {
		return &protopb.OperationConfirmed{Ok: false}, err
	}
	return &protopb.OperationConfirmed{Ok: true}, nil
}

func (rt *router) GetTotalAmount(ctx context.Context, order *protopb.NewOrderCreated) (*protopb.Amount, error) {
	amount := rt.service.GetTotalAmount(ctx, order.OrderID)
	return &protopb.Amount{Amount: float32(amount)}, nil
}

func (rt *router) RemoveAll(ctx context.Context, order *protopb.NewOrderCreated) (*protopb.OperationConfirmed, error) {
	if err := rt.service.DeleteOrder(ctx, order.OrderID); err != nil {
		return &protopb.OperationConfirmed{Ok: false}, err
	}
	return &protopb.OperationConfirmed{Ok: true}, nil
}
