package services

import (
	"errors"
	"net/http"
	"paygateway/models"
	"paygateway/payment_providers"
)

func (a *Services) Capture(transaction *models.Transaction, amount float64) (int, error) {

	if !transaction.CanCapture() {
		return http.StatusNotAcceptable, errors.New("invalid status to capture transaction")
	}

	cardProvider := payment_providers.NewCardProviderMock()
	err := cardProvider.Capture(transaction.UUID.String())
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}
	if amount <= transaction.Amount {
		transaction.Status = "capture"
		transaction.Amount -= amount
		transaction.Spent += amount
	} else {
		return http.StatusForbidden, errors.New("can't capture more than you authorized")
	}

	return http.StatusOK, nil
}
