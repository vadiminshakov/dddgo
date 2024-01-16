package repository

import (
	"database/sql"
)

type outboxRepo struct {
	db *sql.DB
}

func NewOutboxRepo(db *sql.DB) *outboxRepo {
	return &outboxRepo{db: db}
}

func (r *outboxRepo) Save(key string, value []byte) error {
	return nil
}
