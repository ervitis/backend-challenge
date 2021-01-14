package domain

import "github.com/ervitis/backend-challenge/clientrest/repository"

type (
	basket struct {
		repository repository.IRepository
	}

	IBasket interface {
		CreateOrder(int) int
	}
)

func NewBasketService() IBasket {
	return &basket{repository: repository.NewRepository()}
}

func (b *basket) CreateOrder(userID int) int {
	order := Order{
		UserID:   userID,
		Products: make([]Product, 0),
	}

	return b.repository.Save(order)
}

func (b *basket) UpdateOrder(order Order) error {
	return b.repository.Update(order)
}

func (b *basket) DeleteOrder(orderID int) error {
	return b.repository.Delete(orderID)
}

func (b *basket) GetOrderBy(field string, data interface{}) *Order {
	return b.repository.GetBy(field, data)
}
