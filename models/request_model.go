package models

import "github.com/google/uuid"

// /login

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// /authorise

type AuthoriseRequest struct {
	Card   *CardDetails   `json:"card"`
	Amount *AmountDetails `json:"amount"`
}

type CardDetails struct {
	CardNumber      string `json:"card_number"`
	ExpireMonthYear string `json:"expire_monthday"`
	Cvv             int    `json:"cvv"`
}

func (c *CardDetails) GetExpireMonth() string {
	return c.ExpireMonthYear[:2]
}

func (c *CardDetails) GetExpireYear() string {
	return c.ExpireMonthYear[2:]
}

type AmountDetails struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// /capture

type CaptureRequest struct {
	TransactionId *uuid.UUID `json:"transaction_id"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
}

// /void

type VoidRequest struct {
	TransactionId *uuid.UUID `json:"transaction_id"`
}

// /refund

type RefundRequest struct {
	TransactionId *uuid.UUID `json:"transaction_id"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
}
