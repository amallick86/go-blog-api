package routes

import (
	"go-blog-api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		Url:          "/posts",
		Method:       http.MethodGet,
		Handler:      controllers.GetPosts,
		AuthRequired: false,
	},
	{
		Url:          "/posts",
		Method:       http.MethodPost,
		Handler:      controllers.CreatePost,
		AuthRequired: true,
	},
	{
		Url:          "/posts/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetPost,
		AuthRequired: false,
	},
	{
		Url:          "/posts/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		Url:          "/posts/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeletePost,
		AuthRequired: true,
	},
}
