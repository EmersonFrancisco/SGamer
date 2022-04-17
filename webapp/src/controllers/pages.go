package controllers

import (
	"net/http"
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
	utils.ExecuteTemplate(w, "home.html", nil)
}
