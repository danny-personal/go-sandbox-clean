package repositories

import "github.com/danny-personal/go-sandbox-clean/internal/domain/entities"

type PaymentRepository interface {
	GetPaymentID(limit int) (*[]entities.Payment, error)
}
