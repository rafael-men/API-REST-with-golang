package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLStorage cria uma nova conexão MySQL com base na string DSN fornecida
func NewMySQLStorage(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
