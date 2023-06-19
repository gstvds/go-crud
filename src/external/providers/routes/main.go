package routes

import "github.com/gofiber/fiber/v2"

func NewRouter() *fiber.App {
	app := New()
	return Configure(app)
}
