package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config utiliza as variáveis de ambiente para a criação do SecureCookie
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save registra no browser as informações de autenticação
func Save(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encodedDate, erro := s.Encode("data", data)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "SecurityData",
		Value:    encodedDate,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read retorna os valores armazenados no Cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("data")
	if erro != nil {
		return nil, erro
	}

	values := make(map[string]string)
	if erro = s.Decode("data", cookie.Value, &values); erro != nil {
		return nil, erro
	}
	return values, nil
}
