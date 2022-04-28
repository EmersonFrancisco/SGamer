package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisitions"
	"webapp/src/response"

	"github.com/gorilla/mux"
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

// FollowUser chama a API para registrar que o usuário seguiu outro usuário
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/user/%d/follow", config.ApiUrl, userId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPost, url, nil)
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

// NewUser chama a API para registrar novo usuario do BD
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/user/%d/unfollow", config.ApiUrl, userId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPost, url, nil)
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
