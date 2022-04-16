package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represtenta um usuario que utilizara a rede social
type User struct {
	ID         uint64    `json:"id,omitempty"` // omitempty não repassa informação para json quando nulo
	Username   string    `json:"username,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Pass       string    `json:"pass,omitempty"`
	CreateDate time.Time `json:"CreateDate,omitempty"`
}

// Prepare vai chamar as funções para validade e formatar o usuario recebido
func (user *User) Prepare(stage string) error {

	if erro := user.validated(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validated(stage string) error {

	switch {
	case user.Username == "":
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	case user.Nick == "":
		return errors.New("O Nick é obrigatorio e não pode estar em branco")
	case user.Email == "":
		return errors.New("O Email é obrigatorio e não pode estar em branco")
	}
	if stage == "register" {
		if user.Pass == "" {
			return errors.New("O Pass é obrigatorio e não pode estar em branco")
		}
	}
	return nil
}

func (user *User) format(stage string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
	// validação do formato de e-mail
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}
	if stage == "register" {
		passHash, erro := security.Hash(user.Pass)
		if erro != nil {
			return erro
		}

		user.Pass = string(passHash)
	}
	return nil
}
