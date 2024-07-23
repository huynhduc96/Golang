package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// NewDatabase creates a new database connection
func NewDatabase(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
