package database

import (
	"database/sql"

	"github.com/xandreafonso/gogo/internal/domain/entity"
)

type OrderRepositorySQL struct {
	Db *sql.DB
}

func (r *OrderRepositorySQL) Save(order *entity.Order) error {
	_, err := r.Db.Exec("Insert into ordr (id, price, tax, final_price) values (?, ?, ?, ?)", order.Id, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositorySQL) GetTotalTransactions() (int, error) {
	var total int

	err := r.Db.QueryRow("select count(*) from ordr").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func NewOrderRepositorySQL(db *sql.DB) *OrderRepositorySQL {
	return &OrderRepositorySQL{
		Db: db,
	}
}
