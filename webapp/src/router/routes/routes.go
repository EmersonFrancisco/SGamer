package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, routeHome)

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Authentication(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
