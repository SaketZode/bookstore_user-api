package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
	Error      string `json:"error"`
}

func NewBadRequestError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}
