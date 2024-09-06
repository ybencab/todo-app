package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) *PostgresStore {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	return &PostgresStore{db}
}
