package repository

import (
	"context"
	"database/sql"
)

type tximpl struct {
	tx *sql.Tx
}

type TxManagerRepository struct {
	db *sql.DB
}

func NewTxManagerRepo(db *sql.DB) *TxManagerRepository {
	return &TxManagerRepository{db: db}
}

func (r *TxManagerRepository) Begin(ctx context.Context) (*tximpl, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &tximpl{tx}, nil
}

func (r *tximpl) Commit() error {
	return r.tx.Commit()
}

func (r *TxManagerRepository) WithTx(tx Tx) *RepoRegistry {
	return &RepoRegistry{db: r.db, tx: tx}
}
