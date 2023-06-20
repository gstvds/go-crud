package responses

import (
	"go-crud/src/shared"

	"github.com/gofiber/fiber/v2"
)

func JSON(context *fiber.Ctx, statusCode int, data interface{}) {
	context.Set("Content-Type", "application/json")
	context.Status(statusCode)

	context.JSON(data)
}

// Error returns a JSON formatted error to the client.
func Error(context *fiber.Ctx, statusCode int, serverError shared.ServerError) {
	JSON(context, statusCode, struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Error   string `json:"error"`
	}{
		Message: serverError.Message,
		Code:    serverError.Code,
		Error:   serverError.Error.Error(),
	})
}
