package interfaces

import (
	"database/sql"
	"fmt"
)

type Database interface {
	QueryRow(query string, args ...interface{}) (*sql.Rows, error)
}

type PaymentService struct {
	db Database
}

func NewPaymentService(db Database) *PaymentService {
	return &PaymentService{db}
}

type Payment struct {
	PaymentId int `json:"PaymentId"`
}

func (p *PaymentService) GetPaymentId(limit int) (*[]Payment, error) {
	fmt.Println("GetPaymentId start")
	args := []interface{}{limit}
	rows, err := p.db.QueryRow("select payment_id from payment limit $1", args...)
	if err != nil {
		return nil, err
	}
	var payments []Payment
	fmt.Println("rows next")
	for rows.Next() {
		var paymentID int
		// 行の値を変数に割り当てる、変数からポインタを抽出する
		rows.Scan(&paymentID)
		fmt.Printf("payment_id=%v\n", paymentID)
		payment := Payment{PaymentId: paymentID}
		payments = append(payments, payment)
	}
	return &payments, nil
}
