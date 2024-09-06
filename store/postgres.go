package store

import "database/sql"

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() *PostgreStore {
	return &PostgreStore{}
}
