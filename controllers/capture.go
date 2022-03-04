package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"paygateway/models"
)

func (c *Controller) Capture(w http.ResponseWriter, r *http.Request) {
	captureRequest, err := c.getCaptureRequest(r)
	if err != nil {
		log.Println(err)
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	transaction, err := c.Dao.GetTransaction(captureRequest.TransactionId)
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

	err = c.handleCaptureTestCard(transaction.CardNumber)
	if err != nil {
		log.Println(err)
		ApiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	code, err := c.Services.Capture(transaction, captureRequest.Amount)
	if err != nil {
		log.Println(err)
		ApiError(w, err.Error(), code)
		return
	}

	err = c.Dao.UpdateTransaction(transaction)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
	}

	ReturnJSON(w, models.NewRefundResponse("success", transaction.Amount, transaction.Currency), http.StatusOK)
}

func (c *Controller) getCaptureRequest(r *http.Request) (*models.CaptureRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var captureRequest *models.CaptureRequest

	err = json.Unmarshal(body, &captureRequest)
	if err != nil {
		return nil, err
	}

	return captureRequest, nil
}

func (c *Controller) handleCaptureTestCard(card string) error {
	if card == "4000000000000259" {
		return errors.New("capture failure")
	}
	return nil
}
