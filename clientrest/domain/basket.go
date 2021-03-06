package domain

import (
	"github.com/ervitis/backend-challenge/clientrest/model"
	"github.com/ervitis/backend-challenge/clientrest/rpccaller"
)

type (
	basket struct {
		cli rpccaller.RpcCaller
	}

	IBasket interface {
		CreateOrder(int) (int, error)
		DeleteOrder(int) error
		AddItemToBasket(int, []model.Product) error
		Checkout(int) error
		GetTotalAmount(orderID int) float32
	}
)

func NewBasketService(rpcClient rpccaller.RpcCaller) IBasket {
	return &basket{cli: rpcClient}
}

func (b *basket) CreateOrder(userID int) (int, error) {
	return b.cli.CreateOrder(userID)
}

func (b *basket) DeleteOrder(orderID int) error {
	return b.cli.RemoveAll(orderID)
}

func (b *basket) AddItemToBasket(orderID int, items []model.Product) error {
	return b.cli.AddToBasket(orderID, items)
}

func (b *basket) GetTotalAmount(orderID int) float32 {
	amount, err := b.cli.GetTotalAmount(orderID)
	if err != nil {

	}
	return amount
}

func (b *basket) Checkout(orderID int) error {
	return b.cli.CheckOutOrder(orderID)
}
