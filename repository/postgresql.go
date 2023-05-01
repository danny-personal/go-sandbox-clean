package repository

import (
	"database/sql"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDatabase() (*PostgresDatabase, error) {
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
	return d.DB.Query(query, args...)
}
