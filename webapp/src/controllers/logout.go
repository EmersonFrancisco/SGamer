package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout remove os dados de autenticação do Browser do usuário
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
