package controllers

import "net/http"

// NewUser cria novo usuário no BD
func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando novo usuário!"))
}

// SeachAllUsers busca todos os usuários registrados no BD
func SeachAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários!"))
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
