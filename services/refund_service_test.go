package services

import (
	"paygateway/models"
	"testing"
)

func TestRefund(t *testing.T) {
	transaction := models.Transaction{
		Status:   "capture",
		Amount:   950,
		Currency: "PLN",
		Spent:    50,
	}

	services := Services{}

	services.Refund(&transaction, 50)

	actual := transaction.Amount
	expected := float64(1000)

	if actual != expected {
		t.Errorf("Expected %f but got %f", expected, actual)
	}
}
