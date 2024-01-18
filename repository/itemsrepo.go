package repository

import (
	"examplegood/core/domain/vos"
	"github.com/jackc/pgx/v5"
)

type Items struct {
	db *pgx.Conn
}

func NewItemsRepo(db *pgx.Conn) *Items {
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
