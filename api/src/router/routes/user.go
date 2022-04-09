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
		Funcion:               controllers.SearchAllUsers,
		RequiresAutentication: false,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodGet,
		Funcion:               controllers.SearchUser,
		RequiresAutentication: false,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodPut,
		Funcion:               controllers.UpdateUser,
		RequiresAutentication: false,
	},
	{
		Uri:                   "/user/{userID}",
		Method:                http.MethodDelete,
		Funcion:               controllers.DeleteUser,
		RequiresAutentication: false,
	},
}
