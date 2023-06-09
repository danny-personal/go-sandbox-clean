package datastore

import (
	"database/sql"
	"fmt"

	"github.com/danny-personal/go-sandbox-clean/internal/domain/entities"
	"github.com/danny-personal/go-sandbox-clean/internal/domain/repositories"
)

type Database interface {
	QueryRow(query string, arts ...interface{}) (*sql.Rows, error)
}

type paymentRepository struct {
	db Database
}

type Wrapper struct {
	*sql.DB
}

func (db Wrapper) QueryRow(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func NewPaymentRepository(db *sql.DB) repositories.PaymentRepository {
	return &paymentRepository{db: Wrapper{db}}
}

func (r *paymentRepository) GetPaymentID(limit int) (*[]entities.Payment, error) {
	args := []interface{}{limit}
	rows, err := r.db.QueryRow("select payment_id from payment limit $1", args...)
	if err != nil {
		return nil, err
	}
	var payments []entities.Payment
	for rows.Next() {
		var paymentID int
		// 行の値を変数に割り当てる、変数からポインタを抽出する
		rows.Scan(&paymentID)
		fmt.Printf("payment_id=%v\n", paymentID)

		//payments = append(payments)
		payments = append(payments, entities.Payment{PaymentID: paymentID})
	}
	return &payments, nil
}
