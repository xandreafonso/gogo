package usecase

import (
	"github.com/xandreafonso/gogo/internal/domain/entity"
	"github.com/xandreafonso/gogo/internal/domain/repository"
)

type OrderInput struct {
	Id    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository repository.OrderRepository
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		Id:         order.Id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

func NewCalculateFinalPrice(repository repository.OrderRepository) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: repository,
	}
}
