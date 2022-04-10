package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoute = []Route{
	{
		Uri:                   "/user",
		Method:                http.MethodPost,
		Funcion:               controllers.NewUser,
		RequiresAutentication: false,
	},
	{
		Uri:                   "/user",
		Method:                http.MethodGet,
		Funcion:               controllers.SearchFilterUsers,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodGet,
		Funcion:               controllers.SearchUser,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodPut,
		Funcion:               controllers.UpdateUser,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodDelete,
		Funcion:               controllers.DeleteUser,
		RequiresAutentication: true,
	},
}
