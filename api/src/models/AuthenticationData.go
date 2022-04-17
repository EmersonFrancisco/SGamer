package models

// AutenticationData contem o token e o ID do usuario autenticado para envio por request
type AutenticationData struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
