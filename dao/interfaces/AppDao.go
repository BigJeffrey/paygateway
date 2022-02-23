package interfaces

import (
	"paygateway/models"

	"github.com/google/uuid"
)

type AppDao interface {
	Disconnect()
	GetCards() ([]models.CreditCardData, error)
	UpdateCard(models.CreditCardData) error
	GetTransaction(transactionId *uuid.UUID) (*models.Transaction, error)
	GetMerchantByLogin(login string) (*models.Merchant, error)
	CreateTransaction(transaction *models.Transaction) (uuid.UUID, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(transaction models.Transaction) error
}
