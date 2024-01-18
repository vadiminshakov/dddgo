package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Tx struct {
	db *pgx.Conn
}

type TxManagerRepository struct {
	db *pgx.Conn
}

func NewTxManagerRepo(db *pgx.Conn) *Basket {
	return &Basket{db: db}
}

func (r *TxManagerRepository) Begin(ctx context.Context) (*Tx, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{db: tx.Conn()}, nil
}
