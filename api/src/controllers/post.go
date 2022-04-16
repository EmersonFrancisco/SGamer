package controllers

import (
	"api/src/autentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NewPost cria uma nova publicação
func NewPost(w http.ResponseWriter, r *http.Request) {
	idUser, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	var post models.Post
	if erro = json.Unmarshal(bodyRequest, &post); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	post.AuthorId = idUser
	if erro = post.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	repository := repositories.NewRepositoryPost(db)
	post.Id, erro = repository.Create(post)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusCreated, post)
}

// SearchAllPost busca todas as publicações do usuario logado, e os que ele segue
func SearchFeedPost(w http.ResponseWriter, r *http.Request) {
	userId, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	post, erro := repository.SearchFeed(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, post)
}

// SearchPost busca publicações especificas de acordo com ID
func SearchPost(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameter["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	post, erro := repository.SearchId(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, post)
}

// SearchUserPost busca publicações de um usário de acordo com seu ID
func SearchUserPost(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	post, erro := repository.SearchByUser(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, post)
}

// UpdatePost atualiza publicação selecionada
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	parameter := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameter["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	postDB, erro := repository.SearchId(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postDB.AuthorId != userId {
		response.Erro(w, http.StatusForbidden, errors.New("Não é permitido alterações na publicação de outro usuário!"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	var postAlter models.Post
	if erro = json.Unmarshal(bodyRequest, &postAlter); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	if erro := postAlter.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro := repository.Update(postAlter, postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePost deleta uma publicação
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	parameter := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameter["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	postDB, erro := repository.SearchId(postId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postDB.AuthorId != userId {
		response.Erro(w, http.StatusForbidden, errors.New("Não é permitido exclusão da publicação de outro usuário!"))
		return
	}

	if erro := repository.Delete(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameter["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	if erro = repository.Like(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameter["postId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryPost(db)
	if erro = repository.Unlike(postId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}
