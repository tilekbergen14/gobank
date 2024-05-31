package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountById(int) error
}

type PostgresSql struct {
	db *sql.DB
}

func newPostgresStorage() (*PostgresSql, error) {
	connStr := "user=postgres dbname=postgres password=tqbank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresSql{
		db: db,
	}, nil
}
