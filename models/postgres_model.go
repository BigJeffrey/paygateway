package models

import "github.com/google/uuid"

type CreditCardData struct {
	ID             uuid.UUID
	CardNumber     string
	ExpireMonthDay string
	Cvv            int
	Amount         float64
	Currency       string
}

type Transaction struct {
	UUID                  uuid.UUID
	Status                string
	PaymentMethod         string
	CardNumber            string
	ExpireMonthDay        string
	Cvv                   int
	Amount                float64
	Currency              string
	Spent                 float64
	ProviderTransactionId *string
}

func NewTransaction() *Transaction {
	return &Transaction{
		UUID:   uuid.New(),
		Status: "authorized",
	}
}

func (t Transaction) CanVoid() bool {
	return t.Status == "authorized"
}

func (t Transaction) CanCapture() bool {
	return t.Status == "authorized" || t.Status == "capture"
}

func (t Transaction) CanRefund() bool {
	return t.Status == "capture" || t.Status == "refunded"
}

type Merchant struct {
	UUID     uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
}
