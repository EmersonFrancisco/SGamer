package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPost = []Route{
	{
		Uri:                   "/post",
		Method:                http.MethodPost,
		Funcion:               controllers.NewPost,
		RequiresAutentication: true,
	},

	{
		Uri:                   "/post",
		Method:                http.MethodGet,
		Funcion:               controllers.SearchAllPost,
		RequiresAutentication: true,
	},

	{
		Uri:                   "/post/{postId}",
		Method:                http.MethodGet,
		Funcion:               controllers.SearchPost,
		RequiresAutentication: true,
	},

	{
		Uri:                   "/post/{postId}",
		Method:                http.MethodPut,
		Funcion:               controllers.UpdatePost,
		RequiresAutentication: true,
	},

	{
		Uri:                   "/post/{postId}",
		Method:                http.MethodDelete,
		Funcion:               controllers.DeletePost,
		RequiresAutentication: true,
	},
}
