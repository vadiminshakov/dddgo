package antipatterns

import (
	"database/sql"

	"github.com/vadiminshakov/dddgo/core/domain/vos"
	"github.com/vadiminshakov/dddgo/repository"
	"github.com/vadiminshakov/dddgo/repository/queries"
)

type TxItems struct {
	db *sql.DB
	tx repository.Tx
}

func NewItemsRepo(db *sql.DB, tx repository.Tx) *TxItems {
	if tx == nil {
		return &TxItems{tx: db}
	}
	return &TxItems{tx: tx}
}

func (r *TxItems) GetByBasketID(tx *sql.Tx, id int64) (*vos.BasketItem, error) {
	// TODO: implement

	return &vos.BasketItem{}, nil
}

func (r *TxItems) Save(tx *sql.Tx, item *vos.BasketItem) error {
	_, err := tx.Exec(queries.ItemsSave, item.BasketID, item.GoodID, item.Quantity, item.Price, item.Weight)

	return err
}
