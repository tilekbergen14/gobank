package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
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

func (s *PostgresSql) init() error {
	query := `create table if not exists account (
		id serial primary key,
		firstName varchar(50),
		secondName varchar(50),
		number serial,
		balance serial,
		createdAt timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresSql) GetAccounts() ([]*Account, error) {
	res, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for res.Next() {
		account := &Account{}
		if err := res.Scan(
			&account.ID,
			&account.FirstName,
			&account.SecondName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgresSql) CreateAccount(acc *Account) error {
	query := `
			insert into account (firstName, secondName, number, balance, createdAt)
			values ($1, $2, $3, $4, $5)
	`
	_, err := s.db.Exec(query, acc.FirstName, acc.SecondName, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresSql) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresSql) DeleteAccount(int) error {
	return nil
}
func (s *PostgresSql) GetAccountById(int) (*Account, error) {
	return nil, nil
}
