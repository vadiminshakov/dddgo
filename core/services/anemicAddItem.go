// WARN! Example of workflow with anemic domain model. Domain logic is in service layer. Don't do this.
package services

import (
	"errors"
	"examplegood/broker"
	"examplegood/core/domain/vos"
)

var (
	ErrFullBasket = errors.New("basket is full")
)

const (
	maxBasketItems  = 30
	maxBasketWeight = 10000 // 10 кг
)

type anemicBasketService struct {
	basketRepo BasketRepository
	producer   broker.Producer
}

func NewAnemicBasketService(basketRepo BasketRepository) *anemicBasketService {
	return &anemicBasketService{basketRepo: basketRepo}
}

func (s *anemicBasketService) AddItem(basketID int64, item vos.BasketItem) error {
	basket, err := s.basketRepo.GetByID(basketID)
	if err != nil {
		return err
	}

	if len(basket.Items) == maxBasketItems {
		return ErrFullBasket
	}

	if basket.TotalWeight+item.Weight > maxBasketWeight {
		return ErrFullBasket
	}

	for i, itemFromBasket := range basket.Items {
		if itemFromBasket.GoodID == item.GoodID {
			basket.Items[i].Quantity += item.Quantity
		}
	}

	basket.Items = append(basket.Items, vos.BasketItem{GoodID: item.GoodID, Quantity: item.Quantity})
	basket.TotalWeight += item.Weight

	if err := s.basketRepo.Save(basket); err != nil {
		return err
	}

	for _, event := range basket.Events {
		if err := s.producer.Produce(event.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
