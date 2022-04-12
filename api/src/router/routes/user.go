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
	{
		Uri:                   "/user/{userID}/follow", // user id é o id do usuario a ser seguido
		Method:                http.MethodPost,
		Funcion:               controllers.FollowUser,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}/unfollow", // user id é o id do usuario a ser seguido
		Method:                http.MethodPost,
		Funcion:               controllers.UnfollowUser,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}/followers", // user id é o id do usuario que é seguido
		Method:                http.MethodGet,
		Funcion:               controllers.SearchFollowers,
		RequiresAutentication: true,
	},
	{
		Uri:                   "/user/{userID}/following", // user id é o id do usuario que segue
		Method:                http.MethodGet,
		Funcion:               controllers.SearchFollowing,
		RequiresAutentication: true,
	},
}
