package cachestore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"test/internal/model"
	"test/internal/store/cachestore/postgre"
)

type Store struct {
	Repository map[string]*model.Order
	PGStore    *postgre.SQLStore
	sync.Mutex
}

func (s *Store) Create(order *model.Order) error {
	s.Repository[order.OrderUid] = order

	if err := s.PGStore.Create(order); err != nil {
		return err
	}
	return nil
}

func (s *Store) FindByOrderUID(orderUID string) (*model.Order, error) {
	u, ok := s.Repository[orderUID]
	if !ok {
		return nil, fmt.Errorf("order with ID: %v is not exist", orderUID)
	}
	return u, nil
}

func NewCacheStore(driverName, dataSourceName string) (*Store, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	pqStore := postgre.New(db)
	rep, err := pqStore.GetAll()
	if err != nil {
		return nil, err
	}

	return &Store{rep, pqStore, sync.Mutex{}}, nil
}
