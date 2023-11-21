package repository

import "github.com/xandreafonso/gogo/internal/domain/entity"

type OrderRepository interface {
	Save(order *entity.Order) error
	GetTotalTransactions() (int, error)
}
