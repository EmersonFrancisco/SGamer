package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeLogout = Route{

	Uri:                    "/logout",
	Method:                 http.MethodGet,
	Function:               controllers.Logout,
	RequiresAuthentication: true,
}
