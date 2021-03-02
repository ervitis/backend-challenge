package repository

import (
	"github.com/ervitis/backend-challenge/basket/model"
	"reflect"
	"sort"
)

type (
	IRepository interface {
		Save(order model.Order) int
		Delete(orderID int) error
		Update(order model.Order) error
		GetBy(name string, data interface{}) *model.Order
		AddItem(orderID int, item []model.Product)
		RemoveItem(orderID int, itemID string)
	}

	repository struct {
		data    []model.Order
		ordered bool
	}
)

func NewRepository() IRepository {
	return &repository{data: make([]model.Order, 0)}
}

func (repo *repository) getLastID() int {
	if len(repo.data) == 0 {
		return 0
	}

	if repo.ordered {
		return repo.data[0].ID
	}

	sort.Slice(repo.data, func(i, j int) bool {
		return repo.data[i].ID > repo.data[j].ID
	})

	repo.ordered = true

	return repo.data[0].ID
}

func (repo *repository) Save(order model.Order) int {
	lastID := repo.getLastID() + 1

	order.ID = lastID

	repo.data = append(repo.data, order)
	repo.ordered = false

	return order.ID
}

func (repo *repository) Delete(orderID int) error {
	for i, v := range repo.data {
		if v.ID == orderID {
			repo.data = repo.removeOrder(repo.data, i)
		}
	}
	return nil
}

func (repo *repository) Update(order model.Order) error {
	for k, v := range repo.data {
		if v.ID == order.ID {
			repo.data[k] = order
		}
	}

	return nil
}

func (repo *repository) GetBy(name string, data interface{}) *model.Order {
	for _, v := range repo.data {
		r := reflect.ValueOf(v).Elem()
		ty := r.Type()
		for i := 0; i < r.NumField(); i++ {
			if ty.Field(i).Name == name && r.Field(i).Interface() == data {
				return &v
			}
		}
	}

	return nil
}

func (repo *repository) AddItem(orderID int, item []model.Product) {
	for k, v := range repo.data {
		if v.ID == orderID {
			repo.data[k].Products = append(repo.data[k].Products, item...)
		}
	}
}

func (repo *repository) RemoveItem(orderID int, itemID string) {
	var products []model.Product

	for _, v := range repo.data {
		if v.ID == orderID {
			products = v.Products
			break
		}
	}

	if products == nil {
		return
	}

	for k, v := range products {
		if v.ProductID == itemID {
			products = repo.removeItems(products, k)
		}
	}
}

func (repo *repository) removeItems(data []model.Product, i int) []model.Product {
	data[i] = data[len(data)-1]
	return data[:len(data)-1]
}

func (repo *repository) removeOrder(data []model.Order, i int) []model.Order {
	data[i] = data[len(data)-1]
	return data[:len(data)-1]
}
