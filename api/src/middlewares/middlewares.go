package middlewares

import (
	"api/src/autentication"
	"api/src/response"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// authenticate verifica se o usuario ao fazer a requisição está autenticado
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autentication.ValidateToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}

}
