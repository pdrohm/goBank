package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface{
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int)(*Account, error) 
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=gobank password=gobank sslmode=disable host=localhost port=5431"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
			return nil, err
	}

	if err := db.Ping(); err != nil {
			return nil, err
	}

	return &PostgressStore{
			db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    number BIGINT,
    balance BIGINT,
		created_at timestamp
)`

			_, err := s.db.Exec(query)
			return err


}

func (s *PostgressStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
} 

func (s *PostgressStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
} 
