package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"paygateway/models"
)

func (c *Controller) Void(w http.ResponseWriter, r *http.Request) {
	voidRequest, err := c.getVoidRequest(r)
	if err != nil {
		log.Println(err)
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	transaction, err := c.Dao.GetTransaction(voidRequest.TransactionId)
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

	code, err := c.Services.Void(transaction)
	if err != nil {
		log.Println(err)
		ReturnJSON(w, models.NewErrorResponse(err.Error()), code)
		return
	}

	transaction.Status = "void"

	err = c.Dao.UpdateTransaction(transaction)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
	}

	ReturnJSON(w, models.NewVoidResponse("Success", transaction.Amount, transaction.Currency), http.StatusOK)
}

func (c *Controller) getVoidRequest(r *http.Request) (*models.VoidRequest, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var voidRequest *models.VoidRequest

	err = json.Unmarshal(body, &voidRequest)
	if err != nil {
		return nil, err
	}

	return voidRequest, nil
}
