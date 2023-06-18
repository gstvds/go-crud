package responses

import (
	"encoding/json"
	"go-crud/src/shared"
	"log"
	"net/http"
)

// JSON Returns a JSON response to the client.
func JSON(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(response).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// Error returns a JSON formatted error to the client.
func Error(response http.ResponseWriter, statusCode int, serverError shared.ServerError) {
	JSON(response, statusCode, struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Error   string `json:"error"`
	}{
		Message: serverError.Message,
		Code:    serverError.Code,
		Error:   serverError.Error.Error(),
	})
}
