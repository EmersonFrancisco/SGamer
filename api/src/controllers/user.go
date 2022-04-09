package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// NewUser cria novo usuário no BD
func NewUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}
	db, erro := banco.Conect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	userId, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Criando novo usuário com ID: %d", userId)))
}

// SeachAllUsers busca todos os usuários registrados no BD
func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	users, erro := repository.SearchAllUsers()
	if erro != nil {
		log.Fatal(erro)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

// SearchUser busca o usuário registrado no BD com ID informado no parâmetro
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário com ID informado!"))
}

// UptadeUser atualiza os dados do usuário registrado no BD com ID informado no parâmetro
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário com ID informado!"))
}

// DeleteUser deleta o usuário registrado no BD com ID informado no parâmetro
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário com ID informado!"))
}
