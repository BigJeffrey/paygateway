package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"paygateway/models"
	"strconv"

	creditcard "github.com/hubcash/cards" //Basic credit card validation using the Luhn algorithm
)

func (c *Controller) Authorize(w http.ResponseWriter, r *http.Request) {
	authorizeRequest, err := c.getRequest(r)
	if err != nil {
		log.Println(err)
		ApiError(w, "Can not decode request", http.StatusBadRequest)
		return
	}

	err = c.handleAuthorizeTestCard(authorizeRequest.Card.CardNumber)
	if err != nil {
		log.Println(err)
		ApiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = c.validateRequest(authorizeRequest.Card)
	if err != nil {
		log.Println(err)
		ApiError(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	uuid, code, err := c.Services.Authorize(authorizeRequest)
	if err != nil {
		log.Println(err)
		ApiError(w, "Can not authorize payment", code)
		return
	}

	ReturnJSON(w, models.NewAuthoriseResponse(uuid, authorizeRequest.Amount.Amount, authorizeRequest.Amount.Currency), http.StatusOK)
}

func (c *Controller) getRequest(r *http.Request) (*models.AuthoriseRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var authoriseRequest *models.AuthoriseRequest

	err = json.Unmarshal(body, &authoriseRequest)
	if err != nil {
		return nil, err
	}

	return authoriseRequest, nil
}

func (c *Controller) validateRequest(card *models.CardDetails) error {
	validator := creditcard.Card{
		Number:  card.CardNumber,
		Cvv:     strconv.Itoa(card.Cvv),
		Month:   card.GetExpireMonth(),
		Year:    card.GetExpireYear(),
		Company: creditcard.Company{},
	}

	return validator.Validate(true)
}

func (c *Controller) handleAuthorizeTestCard(card string) error {
	if card == "4000000000000119" {
		return errors.New("authorisation failure")
	}
	return nil
}
