package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authentication verifica a existencia de dados nos Cookies
func Authentication(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Read(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		nextFunction(w, r)
	}
}
