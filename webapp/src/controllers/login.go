package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/response"
)

// Login utiliza o e-mail e senha do usuário para autenticar na aplicação
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"pass":  r.FormValue("pass"),
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	resp, erro := http.Post(fmt.Sprintf("%s/login", config.ApiUrl), "application/json", bytes.NewBuffer(user))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}
	var AuthenticationData models.AuthenticationData
	if erro = json.NewDecoder(resp.Body).Decode(&AuthenticationData); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}
	if erro = cookies.Save(w, AuthenticationData.Id, AuthenticationData.Token); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

}
