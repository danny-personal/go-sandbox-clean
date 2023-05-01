package main

import (
	"fmt"
	"log"

	"github.com/danny-personal/go-sandbox-clean/interfaces"
	"github.com/danny-personal/go-sandbox-clean/repository"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDatabase()
	if err != nil {
		log.Fatal()
	}
	defer db.DB.Close()

	paymentService := interfaces.NewPaymentService(db)
	payments, err := paymentService.GetPaymentId(5)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(payments)
}
