package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a reposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna a resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratamentStatusCode trata as requisições com o status code
func TratamentStatusCode(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
