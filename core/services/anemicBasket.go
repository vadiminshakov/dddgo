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

type ItemsRepository interface {
	// GetBasketItems возвращает товары в корзине
	GetByBasketID(id int64) (*vos.BasketItem, error) // Save сохраняет корзину
	// Save сохраняет позицию в корзине
	Save(basket *vos.BasketItem) error
}

type anemicBasketService struct {
	basketRepo BasketRepository
	itemsRepo  ItemsRepository
	producer   broker.Producer
}

func NewAnemicBasketService(basketRepo BasketRepository, itemsRepo ItemsRepository) *anemicBasketService {
	return &anemicBasketService{basketRepo: basketRepo, itemsRepo: itemsRepo}
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

	var itemExists bool
	for i, itemFromBasket := range basket.Items {
		if itemFromBasket.GoodID == item.GoodID {
			itemExists = true
			basket.Items[i].Quantity += item.Quantity
		}
	}
	basket.TotalWeight += item.Weight

	if !itemExists {
		basket.Items = append(basket.Items, &vos.BasketItem{GoodID: item.GoodID, Quantity: item.Quantity})
	}

	for _, itemForSave := range basket.Items {
		if err := s.itemsRepo.Save(itemForSave); err != nil {
			return err
		}
	}

	// oops! we forgot to set basketID to item :(

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
