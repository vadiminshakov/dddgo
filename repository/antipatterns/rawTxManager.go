package antipatterns

import (
	"context"
	"database/sql"
)

type TxManagerRepository struct {
	db *sql.DB
}

func NewTxManagerRepo(db *sql.DB) *TxManagerRepository {
	return &TxManagerRepository{db: db}
}

func (r *TxManagerRepository) Begin(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
