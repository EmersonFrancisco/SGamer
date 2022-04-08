package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("Rodando API na porta:", config.Porta,
		"\n Codigo de conexão do banco:", config.StringConexaoBanco)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}
