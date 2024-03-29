package services

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/vadiminshakov/dddgo/broker"
	"github.com/vadiminshakov/dddgo/core/domain/aggregates"
	"github.com/vadiminshakov/dddgo/core/domain/vos"
	"github.com/vadiminshakov/dddgo/repository"
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

type basketService struct {
	repo     repository.RepositoryRegistry
	producer broker.Producer
}

func NewBasketService(repo repository.RepositoryRegistry) *basketService {
	return &basketService{repo: repo}
}

func (s *basketService) AddItem(basketID int64, item vos.BasketItem) error {
	basket, err := s.repo.Basket().GetByID(basketID)
	if err != nil {
		return err
	}

	if err := basket.AddItem(item); err != nil {
		return err
	}

	if err := s.repo.Basket().Save(basket); err != nil {
		return err
	}

	return s.sendEvents(basket)
}

// not very good naming, but it's for demonstration of transaction method usage
func (s *basketService) AddItemWithTx(ctx context.Context, basketID int64, item vos.BasketItem) error {
	basket, err := s.repo.Basket().GetByID(basketID)
	if err != nil {
		return err
	}

	if err := basket.AddItem(item); err != nil {
		return err
	}

	return s.repo.Transaction(ctx, func(repo repository.RepositoryRegistry) error {
		if err := s.repo.Basket().Save(basket); err != nil {
			return err
		}
		bytesBasket, err := json.Marshal(basket)
		if err != nil {
			return err
		}
		if err := s.repo.Outbox().Save(strconv.Itoa(int(basketID)), bytesBasket); err != nil {
			return err
		}

		return nil
	})
}

func (s *basketService) sendEvents(basket *aggregates.Basket) error {
	for _, event := range basket.Events {
		if err := s.producer.Produce(event.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
