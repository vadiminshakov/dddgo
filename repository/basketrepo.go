package repository

import (
	"examplegood/core/domain/aggregates"
	"github.com/jackc/pgx/v5"
)

type Basket struct {
	db *pgx.Conn
}

func NewBasketRepo(db *pgx.Conn) *Basket {
	return &Basket{db: db}
}

func (r *Basket) GetByID(id int64) (*aggregates.Basket, error) {
	// TODO: implement

	return &aggregates.Basket{}, nil
}

func (r *Basket) Save(basket *aggregates.Basket) error {
	// TODO: implement

	return nil
}
