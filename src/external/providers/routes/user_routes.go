package routes

import (
	"go-crud/src/controllers"
	"net/http"
)

func NewUserRoutes() []Route {
	userController := controllers.NewUserController()

	var routes = []Route{
		{
			URI:      "/user",
			Method:   http.MethodPost,
			Function: userController.Create,
		},
		{
			URI:      "/user/:id",
			Method:   http.MethodGet,
			Function: userController.Get,
		},
		{
			URI:      "/user/:id",
			Method:   http.MethodPut,
			Function: userController.Update,
		},
		{
			URI:      "/user/:id",
			Method:   http.MethodDelete,
			Function: userController.Delete,
		},
	}

	return routes
}
