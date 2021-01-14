package repository

import (
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"reflect"
	"sort"
)

type (
	IRepository interface {
		Save(order domain.Order) int
		Delete(orderID int) error
		Update(order domain.Order) error
		GetBy(name string, data interface{}) *domain.Order
	}

	repository struct {
		data    []domain.Order
		ordered bool
	}
)

func NewRepository() IRepository {
	return &repository{data: make([]domain.Order, 0)}
}

func (repo *repository) getLastID() int {
	if repo.ordered {
		return repo.data[0].ID
	}

	sort.Slice(repo.data, func(i, j int) bool {
		return repo.data[i].ID > repo.data[j].ID
	})

	repo.ordered = true

	return repo.data[0].ID
}

func (repo *repository) Save(order domain.Order) int {
	lastID := repo.getLastID() + 1

	order.ID = lastID

	repo.data = append(repo.data, order)
	repo.ordered = false

	return order.ID
}

func (repo *repository) Delete(orderID int) error {
	for i, v := range repo.data {
		if v.ID == orderID {
			repo.data = append(repo.data[:i], repo.data[i:]...)
		}
	}
	return nil
}

func (repo *repository) Update(order domain.Order) error {
	for k, v := range repo.data {
		if v.ID == order.ID {
			repo.data[k] = order
		}
	}

	return nil
}

func (repo *repository) GetBy(name string, data interface{}) *domain.Order {
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
