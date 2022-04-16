package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeLogin = []Route{
	{
		Uri:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenLogin,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenLogin,
		RequiresAuthentication: false,
	},
}
