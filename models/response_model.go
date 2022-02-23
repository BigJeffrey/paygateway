package models

import "github.com/google/uuid"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{Message: message}
}

// /authorize

type AuthoriseResponse struct {
	TransactionId *uuid.UUID `json:"transaction_id"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
}

func NewAuthoriseResponse(transactionId *uuid.UUID, amount float64, currency string) *AuthoriseResponse {
	return &AuthoriseResponse{TransactionId: transactionId, Amount: amount, Currency: currency}
}

// /capture

type CaptureResponse struct {
	Message  string  `json:"message"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewCaptureResponse(message string) *CaptureResponse {
	return &CaptureResponse{Message: message}
}

// /void

type VoidResponse struct {
	Message  string  `json:"message"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewVoidResponse(message string, amount float64, currency string) *VoidResponse {
	return &VoidResponse{Message: message, Amount: amount, Currency: currency}
}

// /refund

type RefundResponse struct {
	Message  string  `json:"message"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewRefundResponse(message string, amount float64, currency string) *RefundResponse {
	return &RefundResponse{Message: message, Amount: amount, Currency: currency}
}
