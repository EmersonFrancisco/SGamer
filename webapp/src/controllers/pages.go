package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requisitions"
	"webapp/src/response"
	"webapp/src/utils"
)

// LoadScreenLogin vai renderizar a tela de login
func LoadScreenLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadScreenNewUser vai renderizar a tela de cadastro
func LoadScreenRegisterUser(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registeruser.html", nil)
}

// LoadScreenNewUser vai renderizar a tela principal com as publicações
func LoadScreenHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/post", config.ApiUrl)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
	}

	var post []models.Post
	if erro = json.NewDecoder(resp.Body).Decode(&post); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
	}
	fmt.Println(post)
	utils.ExecuteTemplate(w, "home.html", post)
}
