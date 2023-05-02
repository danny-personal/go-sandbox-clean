package web

import (
	"net/http"

	"github.com/danny-personal/go-sandbox-clean/internal/interfaces/controllers"
)

func NewRouter(paymentController *controllers.PaymentController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payment", paymentController.GetPaymentID)
	return mux
}
