package repository

import (
	"context"
	"examplegood/core/domain/aggregates"
	"examplegood/core/domain/vos"
	"github.com/jackc/pgx/v5"
	"os"
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
	Transaction(ctx context.Context, fn func(repo *RepoRegistry) error) error
	TxManager() TxManagerRepository
}

type RepoRegistry struct {
	db *pgx.Conn
}

func New(conn *pgx.Conn) (*RepoRegistry, error) {
	if conn != nil {
		return &RepoRegistry{db: conn}, nil
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	return &RepoRegistry{conn}, nil
}

func (r *RepoRegistry) Basket() BasketRepository {
	return NewBasketRepo(r.db)
}

func (r *RepoRegistry) Items() ItemsRepository {
	return NewItemsRepo(r.db)
}

func (r *RepoRegistry) Transaction(ctx context.Context, fn func(repo *RepoRegistry) error) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	newrepo, err := New(tx.Conn())
	if err != nil {
		return err
	}

	if err := fn(newrepo); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return err
}
