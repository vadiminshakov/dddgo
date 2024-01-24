package repository

import (
	"context"
	"database/sql"
	"examplegood/core/domain/aggregates"
	"examplegood/core/domain/vos"
)

// driven port
type BasketRepository interface {
	// GetByID returns basket by ID
	GetByID(id int64) (*aggregates.Basket, error)
	// Save сохраняет корзину
	Save(basket *aggregates.Basket) error
}

// driven port
type ItemsRepository interface {
	// GetBasketItems returns items in basket
	GetByBasketID(id int64) (*vos.BasketItem, error)
	// Save saves one item in basket
	Save(basket *vos.BasketItem) error
}

type Outbox interface {
	Save(key string, value []byte) error
}

type RepositoryRegistry interface {
	Basket() BasketRepository
	Items() ItemsRepository
	Outbox() Outbox
	Transaction(ctx context.Context, fn func(repo RepositoryRegistry) error) error
	TxManager() TxManagerRepository
}

type RepoRegistry struct {
	db *sql.DB
	tx Tx
}

func New(db *sql.DB, tx Tx) (*RepoRegistry, error) {
	if tx == nil {
		return &RepoRegistry{tx: db, db: db}, nil
	}
	return &RepoRegistry{db: db, tx: tx}, nil
}

func (r *RepoRegistry) Basket() BasketRepository {
	return NewBasketRepo(r.db, r.tx)
}

func (r *RepoRegistry) Items() ItemsRepository {
	return NewItemsRepo(r.db, r.tx)
}

func (r *RepoRegistry) Transaction(ctx context.Context, fn func(repo *RepoRegistry) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	newrepo, err := New(r.db, tx)
	if err != nil {
		return err
	}

	if err := fn(newrepo); err != nil {
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

type Tx interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
