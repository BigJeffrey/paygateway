package interfaces

import (
	"paygateway/models"

	"github.com/google/uuid"
)

type Services interface {
	Authorize(request *models.AuthoriseRequest) (*uuid.UUID, int, error)
	Capture(transaction *models.Transaction, amount float64) (int, error)
	Refund(transaction *models.Transaction, amount float64) (int, error)
	Void(transaction *models.Transaction) (int, error)
}
