package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danny-personal/go-sandbox-clean/internal/domain/repositories"
)

type PaymentController struct {
	repo repositories.PaymentRepository
}

func NewPaymentController(repo repositories.PaymentRepository) *PaymentController {
	return &PaymentController{repo: repo}
}

func (pc *PaymentController) GetPaymentID(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}
	payments, err := pc.repo.GetPaymentID(limit)
	if err != nil {
		http.Error(w, "Error fetching payment", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}
