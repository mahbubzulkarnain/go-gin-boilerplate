package response

import (
	"errors"
	"net/http"
)

var (
	// ErrBadRequest ...
	ErrBadRequest = Failure(http.StatusBadRequest, StatusBadRequest)

	// ErrInternalServerError ...
	ErrInternalServerError = Failure(http.StatusInternalServerError, StatusInternalServerError)

	// ErrGatewayTimeout ...
	ErrGatewayTimeout = Failure(http.StatusGatewayTimeout, StatusGatewayTimeout)
)

// Err ...
func Err(err error) *JSON {
	return ErrInternalServerError.SetError(err)
}

// ErrByStatusCode ...
func ErrByStatusCode(statusCode int) *JSON {
	return ErrInternalServerError.SetError(errors.New(http.StatusText(statusCode)))
}

// Failure ...
func Failure(statusCode int, messageCode string) *JSON {
	res := JSON{
		Status:     FAILURE,
		Message:    StatusText(messageCode),
		StatusCode: statusCode,
	}

	res.Error = errors.New(res.Message.EN)
	return &res
}
