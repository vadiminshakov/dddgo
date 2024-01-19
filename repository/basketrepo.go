package repository

import (
	"database/sql"
	"examplegood/core/domain/aggregates"
)

type Basket struct {
	db *sql.DB
	tx Tx
}

func NewBasketRepo(db *sql.DB) *Basket {
	return &Basket{db: db}
}

func (r *Basket) GetByID(id int64) (*aggregates.Basket, error) {
	// TODO: implement
	r.tx.Exec("INSERT INTO x VALUES (1);")
	return &aggregates.Basket{}, nil
}

func (r *Basket) Save(basket *aggregates.Basket) error {
	// TODO: implement

	return nil
}
