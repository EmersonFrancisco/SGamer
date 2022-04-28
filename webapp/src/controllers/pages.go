package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisitions"
	"webapp/src/response"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadScreenLogin vai renderizar a tela de login
func LoadScreenLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
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
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}

	var post []models.Post
	if erro = json.NewDecoder(resp.Body).Decode(&post); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Post   []models.Post
		UserId uint64
	}{
		Post:   post,
		UserId: userId,
	})
}

// LoadScreenUpdatePost vai renderizar a tela para edição de publicação
func LoadScreenUpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/post/%d", config.ApiUrl, postId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
	}

	var post models.Post
	if erro := json.NewDecoder(resp.Body).Decode(&post); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "updatePost.html", post)
}

// LoadScreenUsers vai renderizar a pagina com usuários referente a pesquisa
func LoadScreenUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/user?user=%s", config.ApiUrl, nameOrNick)

	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}
	var users []models.User
	if erro = json.NewDecoder(resp.Body).Decode(&users); erro != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErroAPI{Erro: erro.Error()})
		return
	}
	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadScreenProfile renderiza a pagina do perfil do usuario
func LoadScreenProfile(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	user, erro := models.SearchCompleteUser(userId, r)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	cookie, _ := cookies.Read(r)
	userLogginId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		UserLogginId uint64
	}{
		User:         user,
		UserLogginId: userLogginId,
	})
}
