package postgre

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"test/internal/model"
)

type SQLStore struct {
	db *sql.DB
}

func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) Create(order *model.Order) error {
	var uID string

	marshal, err := json.Marshal(order)
	if err != nil {
		return err
	}
	if err := s.db.QueryRow("INSERT INTO wborder VALUES ($1, $2) ON CONFLICT (OrderUID) DO UPDATE SET wborder = EXCLUDED.wborder RETURNING OrderUID ", order.OrderUid, marshal).Scan(&uID); err != nil {
		return err
	}
	log.Println("order created")
	return nil
}

func (s *SQLStore) GetAll() (map[string]*model.Order, error) {
	var b []byte
	m := make(map[string]*model.Order)

	rows, err := s.db.Query("SELECT wborder FROM wborder")

	if rows == nil {
		log.Println("Empty store created")
		return m, nil
	}
	for rows.Next() {
		rows.Scan(&b)
		order := &model.Order{}

		if err := json.Unmarshal(b, &order); err != nil {
			return nil, err
		}
		m[order.OrderUid] = order
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	log.Println("Store is ready")
	return m, nil
}
