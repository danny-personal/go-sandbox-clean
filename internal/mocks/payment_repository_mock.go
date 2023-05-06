package mocks

import (
	"github.com/danny-personal/go-sandbox-clean/internal/domain/entities"
	"github.com/danny-personal/go-sandbox-clean/internal/domain/repositories"
)

type PaymentRepositoryMock struct {
	mockGetPayment func(int) (*[]entities.Payment, error)
}

func NewPaymentRepositoryMock() repositories.PaymentRepository {
	return &PaymentRepositoryMock{}
}

func (u *PaymentRepositoryMock) GetPaymentID(limit int) (*[]entities.Payment, error) {
	if u.mockGetPayment != nil {
		return u.mockGetPayment(limit)
	}
	return nil, nil
}

func (u *PaymentRepositoryMock) SetMockGetPaymentID(fn func(int) (*[]entities.Payment, error)) {
	u.mockGetPayment = fn
}
