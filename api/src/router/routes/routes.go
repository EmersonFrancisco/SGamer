package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Funcion).Methods(route.Method)
	}

	return r

}
