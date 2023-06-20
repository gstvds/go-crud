package routes

import "github.com/gofiber/fiber/v2"

type Route struct {
	URI      string
	Method   string
	Function func(*fiber.Ctx) error
}

// Configure a new Fiber router adding routes
func Configure(app *fiber.App) *fiber.App {
	userRoutes := NewUserRoutes()

	for _, userRoute := range userRoutes {
		app.Add(userRoute.Method, userRoute.URI, userRoute.Function)
	}

	return app
}

// NewRouter return a new instance of a Fiber App with routes configured
func NewRouter() *fiber.App {
	app := fiber.New()

	return Configure(app)
}
