package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"paygateway/models"
)

func (c *Controller) Refund(w http.ResponseWriter, r *http.Request) {
	refundRequest, err := c.getRefundRequest(r)
	if err != nil {
		log.Println(err)
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	transaction, err := c.Dao.GetTransaction(refundRequest.TransactionId)
	if err != nil {
		log.Println(err)
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		log.Println(err)
		ApiError(w, "transaction not found", http.StatusNotFound)
		return
	}

	err = c.handleRefundTestCard(transaction.CardNumber)
	if err != nil {
		log.Println(err)
		ApiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	code, err := c.Services.Refund(transaction, refundRequest.Amount)
	if err != nil {
		ApiError(w, err.Error(), code)
		return
	}

	ReturnJSON(w, models.NewRefundResponse("Success", transaction.Amount, transaction.Currency), http.StatusOK)
}

func (c *Controller) getRefundRequest(r *http.Request) (*models.RefundRequest, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var refundRequest *models.RefundRequest

	err = json.Unmarshal(body, &refundRequest)
	if err != nil {
		return nil, err
	}

	return refundRequest, nil
}

func (c *Controller) handleRefundTestCard(card string) error {
	if card == "4000000000003238" {
		return errors.New("refund failure")
	}
	return nil
}
