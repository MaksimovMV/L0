package store

import "test/internal/model"

type Store interface {
	Create(order *model.Order) error
	FindByOrderUID(orderUID string) (*model.Order, error)
}
