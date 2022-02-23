package payment_providers

import "github.com/google/uuid"

type CardProviderInterface interface {
	Authorize(cardNumber string, monthYear string, cvv int, amount float64, currency string) (string, error)
	Capture(transactionId string) error
	Void(transactionId string) error
	Refund(transactionId string) error
}

type CardProviderMock struct {
}

func NewCardProviderMock() *CardProviderMock {
	return &CardProviderMock{}
}

func (c *CardProviderMock) Authorize(cardNumber string, monthYear string, cvv int, amount float64, currency string) (string, error) {
	return uuid.New().String(), nil
}

func (c *CardProviderMock) Capture(transactionId string) error {
	return nil

}

func (c *CardProviderMock) Void(transactionId string) error {
	return nil

}

func (c *CardProviderMock) Refund(transactionId string) error {
	return nil

}
