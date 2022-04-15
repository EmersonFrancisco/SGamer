package controllers

import (
	"api/src/autentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// NewUser lê o request, converte dados e conecta o banco e
// execulta função que cria um novo usuário no BD
func NewUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	if erro = user.Prepare("register"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusCreated, user)
}

// SeachAllUsers lê o filtro informado no request conecta o banco,
// e execulta função que busca os usuários que responde ao filtro no BD
func SearchFilterUsers(w http.ResponseWriter, r *http.Request) {
	filter := strings.ToLower(r.URL.Query().Get("user"))
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	users, erro := repository.SearchFilter(filter)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

// SearchUser lê o id informado por parametro, conecta o banco,
// e execulta função que busca usuario com id no BD
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["userID"], 10, 64)
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
	repository := repositories.NewRepositoryUser(db)
	user, erro := repository.Search(ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// UptadeUser lê o id informado por parametros, conecta o banco,
// e execulta função que atualiza os dados do usuário com id registrado no BD
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["userID"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	idUserToken, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	if ID != idUserToken {
		response.Erro(w, http.StatusForbidden, errors.New("Não permitido atualizar outro usuario, sem ser o seu."))
	}
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
	if erro = user.Prepare("update"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	erro = repository.Update(ID, user)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser lê o id informado por parametro, conecta o banco,
// e execulta função que deleta o usuário registrado com id no BD
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["userID"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	idUserToken, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	if ID != idUserToken {
		response.Erro(w, http.StatusForbidden, errors.New("Não permitido deletar outro usuario, sem ser o seu."))
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	if erro = repository.Delete(ID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// FollowUser lê o ID do usuario logado, e o ID no parametro, conecta no banco
// e execulta funcção que faça que um usuario siga outro
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerUserID, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userID"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if followerUserID == userId {
		response.Erro(w, http.StatusForbidden, errors.New("Não permitido seguir seu próprio usuário."))
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	if erro = repository.Follow(userId, followerUserID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser lê o ID do usuario logado, e o ID no parametro, conecta no banco
// e execulta funcção que faça que um usuario pare de seguir outro
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerUserID, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userID"], 10, 64)
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
	repository := repositories.NewRepositoryUser(db)
	if erro = repository.Unfollow(userId, followerUserID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// SearchFollower lê a ID do parametro, conecta no banco,
// e execulta função que retorna todos usuários que segue esse ID
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userID"], 10, 64)
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
	repository := repositories.NewRepositoryUser(db)
	followers, erro := repository.SearchFollowers(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, followers)
}

// SearchFollower lê a ID do parametro, conecta no banco,
// e execulta função que retorna todos usuários que esse ID segue
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userID"], 10, 64)
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
	repository := repositories.NewRepositoryUser(db)
	users, erro := repository.SearchFollowing(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

//	UpdatePass verifica ID do usuario logado, lê ID do parametro, conecta no banco
// e execulta processor de validação e segurança para alteração da senha do usuário
func UpdatePass(w http.ResponseWriter, r *http.Request) {
	userIdToken, erro := autentication.ExtractIdUser(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
	}
	parameter := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameter["userID"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if userId != userIdToken {
		response.Erro(w, http.StatusForbidden, errors.New("Só é permitida a alteração do seu próprio usuário!"))
		return
	}
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	var pass models.Pass
	if erro = json.Unmarshal(bodyRequest, &pass); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	db, erro := database.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	passDataBase, erro := repository.SearchPass(userId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = security.ValidadePass(pass.Current, passDataBase); erro != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("A senha atual informada é inválida!"))
		return
	}
	passHash, erro := security.Hash(pass.New)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePass(userId, string(passHash)); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
