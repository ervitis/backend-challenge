package domain

import (
	"github.com/ervitis/backend-challenge/clientrest/exceptions"
	"github.com/ervitis/backend-challenge/clientrest/model"
	"github.com/ervitis/backend-challenge/clientrest/repository"
)

type (
	basket struct {
		repository repository.IRepository
	}

	IBasket interface {
		CreateOrder(int) int
		UpdateOrder(order model.Order) error
		DeleteOrder(orderID int) error
		GetOrderBy(field string, data interface{}) *model.Order
		AddItemToBasket(orderID int, items []model.Product) error
		Checkout(orderID int)
		GetTotalAmount(orderID int) int
	}
)

func NewBasketService() IBasket {
	return &basket{repository: repository.NewRepository()}
}

func (b *basket) CreateOrder(userID int) int {
	order := model.Order{
		UserID:   userID,
		Products: make([]model.Product, 0),
	}

	return b.repository.Save(order)
}

func (b *basket) UpdateOrder(order model.Order) error {
	return b.repository.Update(order)
}

func (b *basket) DeleteOrder(orderID int) error {
	return b.repository.Delete(orderID)
}

func (b *basket) GetOrderBy(field string, data interface{}) *model.Order {
	return b.repository.GetBy(field, data)
}

func (b *basket) AddItemToBasket(orderID int, item []model.Product) error {
	if order := b.repository.GetBy("orderID", orderID); order == nil {
		return exceptions.OrderNotFound(orderID)
	}

	b.repository.AddItem(orderID, item)
	return nil
}

func (b *basket) GetTotalAmount(orderID int) int {
	return 1
}

func (b *basket) Checkout(orderID int) {

}
