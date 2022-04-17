package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeHome = Route{
	Uri:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.LoadScreenHome,
	RequiresAuthentication: true,
}
