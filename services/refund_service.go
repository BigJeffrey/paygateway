package services

import (
	"errors"
	"net/http"
	"paygateway/models"
	"paygateway/payment_providers"
)

func (a *Services) Refund(transaction *models.Transaction, amount float64) (int, error) {

	if !transaction.CanRefund() {
		return http.StatusNotAcceptable, errors.New("invalid status to refund transaction")
	}

	cardProvider := payment_providers.NewCardProviderMock()
	err := cardProvider.Refund(transaction.UUID.String())
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}
	if amount <= transaction.Spent {
		transaction.Status = "refunded"
		transaction.Amount += amount
		transaction.Spent -= amount
	} else {
		return http.StatusForbidden, errors.New("can't refund more than you spent")
	}

	err = a.Dao.UpdateTransaction(transaction)
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}

	return http.StatusOK, nil
}
