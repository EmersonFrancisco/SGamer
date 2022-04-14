package models

// Pass representa o formato para alteração de senha
type Pass struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
