package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/danny-personal/go-sandbox-clean/internal/infrastructure/persistence/datastore"
	"github.com/danny-personal/go-sandbox-clean/internal/infrastructure/web"
	"github.com/danny-personal/go-sandbox-clean/internal/interfaces/controllers"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@192.168.0.239/pagila")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()
	paymentRepository := datastore.NewPaymentRepository(db)
	/*
		payment, err := paymentRepository.GetPaymentID(5)
		if err != nil {
			log.Fatal()
		}
		fmt.Println(payment)
	*/
	paymentController := controllers.NewPaymentController(paymentRepository)
	router := web.NewRouter(paymentController)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
