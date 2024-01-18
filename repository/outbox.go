package repository

import (
	"github.com/jackc/pgx/v5"
)

type outboxRepo struct {
	db *pgx.Conn
}

func NewOutboxRepo(db *pgx.Conn) *outboxRepo {
	return &outboxRepo{db: db}
}

func (r *outboxRepo) Save(key string, value []byte) error {
	return nil
}
