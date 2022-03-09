package controllers

import (
	"encoding/json"
	"net/http"

	"paygateway/dao/interfaces"
	"paygateway/models"
)

type Controller struct {
	Dao      interfaces.AppDao
	Services interfaces.Services
}

func ReturnJSON(w http.ResponseWriter, responseModel interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(responseModel)
	if err != nil {
		ReturnJSON(w, models.NewErrorResponse("unable to encode json"), http.StatusInternalServerError)
	}
}

func ApiError(w http.ResponseWriter, message string, code int) {
	ReturnJSON(w, models.NewErrorResponse(message), code)
}

func ApiSuccess(w http.ResponseWriter, message string, code int) {
	ReturnJSON(w, models.NewSuccessResponse(message), code)
}
