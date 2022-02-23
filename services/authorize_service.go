package services

import (
	"errors"
	"log"
	"net/http"
	"paygateway/models"
	"paygateway/payment_providers"

	"github.com/google/uuid"
)

func (a *Services) Authorize(request *models.AuthoriseRequest) (*uuid.UUID, int, error) {
	transaction := models.NewTransaction()
	transaction.PaymentMethod = "credit_card"
	transaction.CardNumber = request.Card.CardNumber
	transaction.ExpireMonthDay = request.Card.ExpireMonthYear
	transaction.Cvv = request.Card.Cvv
	transaction.Amount = request.Amount.Amount
	transaction.Currency = request.Amount.Currency

	uuid, err := a.Dao.CreateTransaction(transaction)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("something went wrong")
	}

	cardProvider := payment_providers.NewCardProviderMock()
	providerTransactionId, err := cardProvider.Authorize(transaction.CardNumber, transaction.ExpireMonthDay, transaction.Cvv, transaction.Amount, transaction.Currency)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("something went wrong")
	}

	transaction.Status = "authorized"
	transaction.ProviderTransactionId = &providerTransactionId
	err = a.Dao.UpdateTransaction(transaction)
	if err != nil {
		log.Println("update transaction", err)
		return nil, http.StatusInternalServerError, errors.New("something went wrong")
	}
	return &uuid, http.StatusOK, nil
}
