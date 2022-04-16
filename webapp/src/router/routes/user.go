package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesUser = []Route{
	{
		Uri:                    "/registeruser",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenRegisterUser,
		RequiresAuthentication: false,
	},
}
