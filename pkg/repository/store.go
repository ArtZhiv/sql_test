package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sqlx.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sqlx.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	s.db = db
	fmt.Println("Open")
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
	fmt.Println("Close")
}
