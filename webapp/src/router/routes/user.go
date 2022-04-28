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
	{
		Uri:                    "/user",
		Method:                 http.MethodPost,
		Function:               controllers.NewUser,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/searchUser",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenUsers,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/user/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenProfile,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/user/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/user/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
}
