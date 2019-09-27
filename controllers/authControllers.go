package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
	u "rest-api/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	acoount := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acoount)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(acoount.Email, acoount.Password)
	u.Respond(w, resp)
}