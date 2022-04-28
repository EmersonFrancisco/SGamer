package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routePost = []Route{
	{
		Uri:                    "/post",
		Method:                 http.MethodPost,
		Function:               controllers.NewPost,
		RequiresAuthentication: true,
	},

	{
		Uri:                    "/post/{postId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},

	{
		Uri:                    "/post/{postId}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/post/{postId}/update",
		Method:                 http.MethodGet,
		Function:               controllers.LoadScreenUpdatePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/post/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/post/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
}
