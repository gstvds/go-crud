package routes

import (
	usercontroller "go-crud/src/controllers/user_controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	URI      string
	Method   string
	Function func(*fiber.Ctx) error
}

var routes = []Route{
	{
		URI:      "/user",
		Method:   http.MethodPost,
		Function: usercontroller.Create,
	},
	{
		URI:      "/user/:id",
		Method:   http.MethodGet,
		Function: usercontroller.Get,
	},
	{
		URI:      "/user/:id",
		Method:   http.MethodPut,
		Function: usercontroller.Update,
	},
	{
		URI:      "/user/:id",
		Method:   http.MethodDelete,
		Function: usercontroller.Delete,
	},
}

func Configure(app *fiber.App) *fiber.App {
	for _, route := range routes {
		app.Add(route.Method, route.URI, route.Function)
	}

	return app
}
