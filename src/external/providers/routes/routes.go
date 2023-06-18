package routes

import (
	usercontroller "go-crud/src/controllers/user_controller"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
	{
		URI:      "/user",
		Method:   http.MethodPost,
		Function: usercontroller.Create,
	},
	{
		URI:      "/user/{id}",
		Method:   http.MethodGet,
		Function: usercontroller.Get,
	},
	{
		URI:      "/user/{id}",
		Method:   http.MethodPut,
		Function: usercontroller.Update,
	},
	{
		URI:      "/user/{id}",
		Method:   http.MethodDelete,
		Function: usercontroller.Delete,
	},
}

func Configure(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
