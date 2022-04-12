package controllers

import (
	"api/src/autentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login é resposavel por efetuar login de um usuario
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repository := repositories.NewRepositoryUser(db)
	userDB, erro := repository.SearchEmail(user.Email)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = security.ValidadePass(user.Pass, userDB.Pass); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	token, erro := autentication.NewToken(userDB.ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
}
