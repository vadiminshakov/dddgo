package repository

import (
	"database/sql"
	"examplegood/core/domain/aggregates"
)

type Basket struct {
	db *sql.DB
}

func NewBasketRepo(db *sql.DB) *Basket {
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
