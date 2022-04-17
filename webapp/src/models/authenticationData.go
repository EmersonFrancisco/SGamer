package models

// AuthenticationData representa os dados que recebera por request para autenticação quando necessario
type AuthenticationData struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
