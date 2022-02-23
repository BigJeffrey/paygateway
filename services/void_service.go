package services

import (
	"errors"
	"net/http"
	"paygateway/models"
	"paygateway/payment_providers"
)

func (a *Services) Void(transaction *models.Transaction) (int, error) {

	if !transaction.CanVoid() {
		return http.StatusNotAcceptable, errors.New("invalid status to void transaction")
	}

	cardProvider := payment_providers.NewCardProviderMock()
	err := cardProvider.Void(transaction.UUID.String())
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}

	transaction.Status = "void"

	err = a.Dao.UpdateTransaction(transaction)
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}

	return http.StatusOK, nil
}
