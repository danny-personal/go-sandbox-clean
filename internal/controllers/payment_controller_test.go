package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danny-personal/go-sandbox-clean/internal/domain/entities"
	"github.com/danny-personal/go-sandbox-clean/internal/interfaces/controllers"
	"github.com/danny-personal/go-sandbox-clean/internal/mocks"
)

func TestPaymentController_GetPayment(t *testing.T) {
	paymentRepo := mocks.NewPaymentRepositoryMock().(*mocks.PaymentRepositoryMock)
	paymentController := controllers.NewPaymentController(paymentRepo)

	paymentRepo.SetMockGetPaymentID(func(limit int) (*[]entities.Payment, error) {
		if limit == 2 {
			payments := []entities.Payment{
				{PaymentID: 1},
				{PaymentID: 2},
			}
			return &payments, nil
		}
		return nil, errors.New("payment not found")
	})

	t.Run("valid limit", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "api/payment?limit=2", nil)

		paymentController.GetPaymentID(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid limit number", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "api/payment?limit=hoge", nil)

		paymentController.GetPaymentID(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}
