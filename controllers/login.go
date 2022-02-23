package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"paygateway/models"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {

	loginRequest, err := c.getLoginRequest(r)
	if err != nil {
		log.Println(err)
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	merchant, err := c.Dao.GetMerchantByLogin(loginRequest.Username)
	if err != nil {
		log.Println(err)
		ApiError(w, "something went wrong", http.StatusBadRequest)
		return
	}

	if merchant == nil || merchant.Password != loginRequest.Password {
		log.Println(err)
		ApiError(w, "incorrect merchant credentials", http.StatusNotFound)
		return
	}

	validToken, err := generateJWT()
	if err != nil {
		log.Println(err)
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   validToken,
			Expires: time.Now().Add(time.Minute * 10),
		})

	ApiSuccess(w, "Merchant signed in", http.StatusAccepted)
}

func (c *Controller) getLoginRequest(r *http.Request) (*models.LoginRequest, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var loginRequest *models.LoginRequest

	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		return nil, err
	}

	return loginRequest, nil
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["autorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySignedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
