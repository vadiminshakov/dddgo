package repository

import (
	"database/sql"
	"examplegood/core/domain/vos"
)

type Items struct {
	db *sql.DB
}

func NewItemsRepo(db *sql.DB) *Items {
	return &Items{db: db}
}

func (r *Items) GetByBasketID(id int64) (*vos.BasketItem, error) {
	// TODO: implement

	return &vos.BasketItem{}, nil
}

func (r *Items) Save(basket *vos.BasketItem) error {
	// TODO: implement

	return nil
}
