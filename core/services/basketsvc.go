package services

import (
	"examplegood/broker"
	"examplegood/core/domain/aggregates"
	"examplegood/core/domain/vos"
)

// driver port
type BasketService interface {
	// AddItem добавляет товар в корзину
	AddItem(basketID int64, goodID int64, quantity int64) error
	// RemoveItem удаляет товар из корзины
	RemoveItem(basketID int64, goodID int64, quantity int64) error
	// TotalPrice возвращает общую стоимость товаров в корзине
	TotalPrice() float64
}

// driven port
type BasketRepository interface {
	// GetByID возвращает корзину по идентификатору
	GetByID(id int64) (*aggregates.Basket, error)
	// Save сохраняет корзину
	Save(basket *aggregates.Basket) error
}

type basketService struct {
	basketRepo BasketRepository
	producer   broker.Producer
}

func NewBasketService(basketRepo BasketRepository) *basketService {
	return &basketService{basketRepo: basketRepo}
}

func (s *basketService) AddItem(basketID int64, item vos.BasketItem) error {
	basket, err := s.basketRepo.GetByID(basketID)
	if err != nil {
		return err
	}

	if err := basket.AddItem(item); err != nil {
		return err
	}

	if err := s.basketRepo.Save(basket); err != nil {
		return err
	}

	return s.sendEvents(basket)
}

func (s *basketService) sendEvents(basket *aggregates.Basket) error {
	for _, event := range basket.Events {
		if err := s.producer.Produce(event.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
