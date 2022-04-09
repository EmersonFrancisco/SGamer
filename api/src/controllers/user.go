package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NewUser lê o request, converte dados e conecta o banco e
// execulta função que cria um novo usuário no BD
func NewUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao ler Request!:%s", erro)))
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter Json para user!:%s", erro)))
		return
	}
	db, erro := banco.Conect()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao conectar com banco!:%s", erro)))
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	userId, erro := repository.Create(user)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro durante a criação do user no Banco!:%s", erro)))
		return
	}
	w.Write([]byte(fmt.Sprintf("Criando novo usuário com ID: %d", userId)))
}

// SeachAllUsers conecta o banco, e execulta função que busca todos os usuários registrados no BD
func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conect()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao conectar ao Banco!:%s", erro)))
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	users, erro := repository.SearchAllUsers()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro durante a busca dos users no Banco:%s", erro)))
		return
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter user para Json!:%s", erro)))
		return
	}
}

// SearchUser lê o id informado por parametro, conecta o banco,
// e execulta função que busca usuario com id no BD
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["id"], 10, 32)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter parametro para inteiro!:%s", erro)))
		return
	}
	db, erro := banco.Conect()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao conecar ao Banco!:%s", erro)))
		return
	}
	repository := repositories.NewRepositoryUser(db)
	user, erro := repository.SearchUser(ID)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro durante a buscar do user no Banco:%s", erro)))
		return
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter user para Json!:%s", erro)))
		return
	}
}

// UptadeUser atualiza os dados do usuário registrado no BD com ID informado no parâmetro
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["id"], 10, 32)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter parametro para inteiro!:%s", erro)))
		return
	}
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao ler Request!:%s", erro)))
		return
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter Json para user!:%s", erro)))
		return
	}
	db, erro := banco.Conect()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao conectar com banco!:%s", erro)))
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	erro = repository.UpdateUser(uint64(ID), user)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao efetuar atualização do user!:%s", erro)))
		return
	}
	w.Write([]byte(fmt.Sprintf("Atualizando usuário com ID %d com sucesso!", ID)))
}

// DeleteUser deleta o usuário registrado no BD com ID informado no parâmetro
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameter["id"], 10, 32)
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao converter parametro para inteiro!:%s", erro)))
		return
	}
	db, erro := banco.Conect()
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao conectar com banco!:%s", erro)))
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUser(db)
	erro = repository.DeleteUser(uint64(ID))
	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro no processo de exclusão do user do banco!:%s", erro)))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Deletado usuário com ID %d com sucesso!", ID)))
}
