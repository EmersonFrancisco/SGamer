package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Println("Rodando Web App na porta: 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
