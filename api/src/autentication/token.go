package autentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// NewToken cria um token com as permissões para o usuario após login
func NewToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// Validate verifica se o token iformado na requisição é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token inválido")
}

// ExtractIdUser retorna o ID do usuario responsavel pelo Token
func ExtractIdUser(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return 0, erro
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return userId, nil
	}

	return 0, errors.New("Token inválido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
