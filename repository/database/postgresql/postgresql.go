package postgresql

import (
	"database/sql"
	"fmt"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase() (*PostgresDatabase, error) {
	fmt.Println("connection start")
	con := "postgres://postgres:password@192.168.0.239/pagila"
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{db}, nil
}

func (d *PostgresDatabase) QueryRow(query string, args ...interface{}) (*sql.Rows, error) {
	//return d.db.QueryRow(query, args...)
	return d.db.Query(query, args...)
}
