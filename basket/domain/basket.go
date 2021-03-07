package domain

import (
	"context"
	"github.com/ervitis/backend-challenge/basket/model"
	"github.com/ervitis/backend-challenge/basket/repository"
	protopb "github.com/ervitis/backend-challenge/proto"
	"github.com/google/uuid"
)

type (
	basket struct {
		repo repository.IRepository
	}

	IBasket interface {
		CreateOrder(context.Context, int64) (int64, error)
		DeleteOrder(context.Context, int64) error
		AddItemToBasket(context.Context, int64, []*protopb.Item) error
		Checkout(context.Context, int64) error
		GetTotalAmount(context.Context, int64) float64
	}
)

func NewBasketService(repo repository.IRepository) IBasket {
	return &basket{repo: repo}
}

func (svc *basket) CreateOrder(ctx context.Context, userID int64) (int64, error) {
	orderByUser := svc.repo.GetBy("userID", userID)
	if orderByUser != nil {
		return orderByUser.ID, nil
	}
	orderByUser = &model.Order{
		ID:     int64(uuid.New().ID()),
		UserID: userID,
	}
	return svc.repo.Save(*orderByUser), nil
}

func (svc *basket) DeleteOrder(ctx context.Context, orderID int64) error {
	order := svc.repo.GetBy("orderID", orderID)
	if order == nil {
		return nil
	}

	return svc.repo.Delete(orderID)
}

func (svc *basket) AddItemToBasket(ctx context.Context, orderID int64, items []*protopb.Item) error {
	var orderItems []model.Product
	for _, item := range items {
		orderItem := model.Product{
			ProductID: item.ProductID,
			Quantity:  int(item.Quantity),
		}
		orderItems = append(orderItems, orderItem)
	}
	svc.repo.AddItem(orderID, orderItems)
	return nil
}

func (svc *basket) Checkout(ctx context.Context, orderID int64) error {
	return nil
}

func (svc *basket) GetTotalAmount(ctx context.Context, orderID int64) float64 {
	order := svc.repo.GetBy("orderID", orderID)
	if order == nil {
		return 0
	}
	amount := 0.0
	for _, item := range order.Products {
		amount += item.Price * float64(item.Quantity)
	}
	return amount
}
