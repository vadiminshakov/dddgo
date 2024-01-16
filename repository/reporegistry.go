package repository

import (
	"database/sql"
	"examplegood/core/domain/aggregates"
	"examplegood/core/domain/vos"
)

// driven port
type BasketRepository interface {
	// GetByID возвращает корзину по идентификатору
	GetByID(id int64) (*aggregates.Basket, error)
	// Save сохраняет корзину
	Save(basket *aggregates.Basket) error
}

// driven port
type ItemsRepository interface {
	// GetBasketItems возвращает товары в корзине
	GetByBasketID(id int64) (*vos.BasketItem, error) // Save сохраняет корзину
	// Save сохраняет позицию в корзине
	Save(basket *vos.BasketItem) error
}

type Outbox interface {
	Save(key string, value []byte) error
}

type Repository interface {
	Basket() BasketRepository
	Items() ItemsRepository
	Outbox() Outbox
	Transaction(fn func(repo *RepoRegistry) error) error
}

type RepoRegistry struct {
	db *sql.DB
}

func New(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{db}
}

func (r *RepoRegistry) Basket() BasketRepository {
	return NewBasketRepo(r.db)
}

func (r *RepoRegistry) Items() ItemsRepository {
	return NewItemsRepo(r.db)
}

func (r *RepoRegistry) Transaction(fn func(repo *RepoRegistry) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := fn(r); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return err
}
