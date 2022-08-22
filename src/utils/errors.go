package utils

type ServerError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Error   error  `json:"error"`
}

// Error returns a new instance of ServerError struct
func Error(message string, code string, err error) ServerError {
	return ServerError{Message: message, Code: code, Error: err}
}
