package repository

import (
	"database/sql"

	"github.com/vadiminshakov/dddgo/core/domain/vos"
	"github.com/vadiminshakov/dddgo/repository/queries"
)

type Items struct {
	db *sql.DB
	tx Tx
}

func NewItemsRepo(db *sql.DB, tx Tx) *Items {
	if tx == nil {
		return &Items{tx: db}
	}
	return &Items{tx: tx}
}

func (r *Items) GetByBasketID(id int64) (*vos.BasketItem, error) {
	// TODO: implement

	return &vos.BasketItem{}, nil
}

func (r *Items) Save(item *vos.BasketItem) error {
	_, err := r.tx.Exec(queries.ItemsSave, item.BasketID, item.GoodID, item.Quantity, item.Price, item.Weight)

	return err
}
