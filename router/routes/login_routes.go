package routes

import (
	"go-blog-api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		Url:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
