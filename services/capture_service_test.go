package services

import (
	"paygateway/models"
	"testing"
)

func TestCapture(t *testing.T) {
	transaction := models.Transaction{
		Status:   "authorized",
		Amount:   1000,
		Currency: "PLN",
		Spent:    50,
	}

	services := Services{}

	_, err := services.Capture(&transaction, 50)
	if err != nil {
		t.Errorf("couldnt capture")
	}
	actual := transaction.Amount
	expected := float64(950)

	if actual != expected {
		t.Errorf("Expected %f but got %f", expected, actual)
	}
}
