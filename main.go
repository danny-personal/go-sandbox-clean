package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello")
	db, err := NewPostgresDatabase()
	if err != nil {
		log.Fatal()
	}
	defer db.db.Close()
	paymentService := NewPaymentService(db)
	payments, err := paymentService.GetPaymentId(5)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(payments)
}

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
