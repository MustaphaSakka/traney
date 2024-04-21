package exception

import "net/http"

type AppException struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NotFoundException(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func InternalServerException(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
