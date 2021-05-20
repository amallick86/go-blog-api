package routes

import (
	"go-blog-api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := usersRoutes
	routes = append(routes, postsRoutes...)
	routes = append(routes, loginRoutes...)
	return routes

}
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.Url,
				middlewares.SetMddlewareLogger(
					middlewares.SetMddlewareJSON(
						middlewares.SetMddlewareAuthentication(route.Handler))),
			).Methods(route.Method)

		} else {
			r.HandleFunc(route.Url,
				middlewares.SetMddlewareLogger(
					middlewares.SetMddlewareJSON(route.Handler)),
			).Methods(route.Method)
		}
	}
	return r
}
