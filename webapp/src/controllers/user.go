package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/response"
)

// NewUser chama a API para registrar novo usuario do BD
func NewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"username": r.FormValue("username"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"pass":     r.FormValue("pass"),
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	resp, erro := http.Post(fmt.Sprintf("%s/user", config.ApiUrl), "application/json", bytes.NewBuffer(user))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}
