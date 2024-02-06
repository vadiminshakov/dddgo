package antipatterns

import (
	"database/sql"

	"github.com/vadiminshakov/dddgo/core/domain/aggregates"
	"github.com/vadiminshakov/dddgo/repository"
	"github.com/vadiminshakov/dddgo/repository/queries"
)

type TxBasket struct {
	db *sql.DB
	tx repository.Tx
}

func NewTxBasketRepo(db *sql.DB, tx repository.Tx) *TxBasket {
	if tx == nil {
		return &TxBasket{tx: db}
	}
	return &TxBasket{tx: tx}
}

func (r *TxBasket) GetByID(tx *sql.Tx, id int64) (*aggregates.Basket, error) {
	tx.Exec(queries.BasketSave)
	return &aggregates.Basket{}, nil
}

func (r *TxBasket) Save(tx *sql.Tx, basket *aggregates.Basket) error {
	_, err := tx.Exec(queries.BasketSave, basket.ID, basket.TotalWeight)

	return err
}
