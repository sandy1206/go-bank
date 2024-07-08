package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccount(id int) (*Account, error)
	DeleteAccount(id int) error
	UpdateAccount(a *Account) error
	GetAccountByID(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

// Init initializes the PostgresStore by creating the necessary database table if it doesn't exist.
// It returns an error if there was a problem executing the SQL statement.
func (s *PostgresStore) Init() error {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS accounts (id serial PRIMARY KEY, first_name TEXT, last_name TEXT, number BIGINT, balance BIGINT)")
	return err
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	_, err := s.db.Exec("INSERT INTO accounts (id, first_name, last_name, number, balance) VALUES ($1, $2, $3, $4, $5)", a.ID, a.FirstName, a.LastName, a.Number, a.Balance)
	return err
}

func (s *PostgresStore) GetAccount(id int) (*Account, error) {
	row := s.db.QueryRow("SELECT id, first_name, last_name, number, balance FROM accounts WHERE id = $1", id)
	a := &Account{}
	err := row.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Number, &a.Balance)
	return a, err
}

func (s *PostgresStore) DeleteAccount(id int) error {
	what, err := s.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	fmt.Println("what", what, err)
	return err
}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	_, err := s.db.Exec("UPDATE accounts SET first_name = $1, last_name = $2, number = $3, balance = $4 WHERE id = $5", a.FirstName, a.LastName, a.Number, a.Balance, a.ID)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	row := s.db.QueryRow("SELECT id, first_name, last_name, number, balance FROM accounts WHERE id = $1", id)
	a := &Account{}
	err := row.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Number, &a.Balance)
	return a, err
}
