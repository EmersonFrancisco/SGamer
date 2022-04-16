package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa uma estrutura de rota
type Route struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Config coloca todas as rotas dentro do router
func Config(router *mux.Router) *mux.Router {
	routes := routeLogin
	routes = append(routes, routesUser...)

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
