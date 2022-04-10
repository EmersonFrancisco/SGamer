package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	Uri                   string
	Method                string
	Funcion               func(http.ResponseWriter, *http.Request)
	RequiresAutentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoute
	routes = append(routes, routeLogin)

	for _, route := range routes {

		if route.RequiresAutentication {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Authenticate(route.Funcion))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Funcion)).Methods(route.Method)
		}
	}

	return r

}
