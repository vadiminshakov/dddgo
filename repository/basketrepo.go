package repository

import (
	"database/sql"
	"examplegood/core/domain/aggregates"
	"examplegood/repository/queries"
)

type Basket struct {
	db *sql.DB
	tx Tx
}

func NewBasketRepo(db *sql.DB, tx Tx) *Basket {
	if tx == nil {
		return &Basket{tx: db}
	}
	return &Basket{tx: tx}
}

func (r *Basket) GetByID(id int64) (*aggregates.Basket, error) {
	r.tx.Exec(queries.BasketSave)
	return &aggregates.Basket{}, nil
}

func (r *Basket) Save(basket *aggregates.Basket) error {
	_, err := r.tx.Exec(queries.BasketSave, basket.ID, basket.TotalWeight)

	return err
}
