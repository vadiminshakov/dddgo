// WARN! Example of workflow with anemic domain model. Domain logic is in service layer. Don't do this.
package antipatterns

import (
	"context"
	"database/sql"

	"github.com/vadiminshakov/dddgo/broker"
	"github.com/vadiminshakov/dddgo/core/domain/aggregates"
	"github.com/vadiminshakov/dddgo/core/domain/vos"
	"github.com/vadiminshakov/dddgo/repository/antipatterns"
)

// driven port
type txBasketRepository interface {
	// GetByID returns basket by ID
	GetByID(tx *sql.Tx, id int64) (*aggregates.Basket, error)
	// Save сохраняет корзину
	Save(tx *sql.Tx, basket *aggregates.Basket) error
}

// driven port
type txItemsRepository interface {
	// GetBasketItems returns items in basket
	GetByBasketID(tx *sql.Tx, id int64) (*vos.BasketItem, error)
	// Save saves one item in basket
	Save(tx *sql.Tx, basket *vos.BasketItem) error
}

type txAnemicBasketService struct {
	basketRepo txBasketRepository
	itemsRepo  txItemsRepository
	producer   broker.Producer
	txManager  antipatterns.TxManagerRepository
}

func NewTxAnemicBasketService(basketRepo txBasketRepository, itemsRepo txItemsRepository) *txAnemicBasketService {
	return &txAnemicBasketService{basketRepo: basketRepo, itemsRepo: itemsRepo}
}

func (s *txAnemicBasketService) AddItem(basketID int64, item vos.BasketItem) error {
	tx, err := s.txManager.Begin(context.Background())
	if err != nil {
		return err
	}

	basket, err := s.basketRepo.GetByID(tx, basketID)
	if err != nil {
		return err
	}

	// some manipulations with basket and item here ...

	for _, itemForSave := range basket.Items {
		if err := s.itemsRepo.Save(tx, itemForSave); err != nil {
			return err
		}
	}

	if err := s.basketRepo.Save(tx, basket); err != nil {
		return err
	}

	return tx.Commit()
}
