package response

import (
	"errors"
	"net/http"
)

var (
	// ErrBadRequest godoc.
	ErrBadRequest = Failure(http.StatusBadRequest, StatusBadRequest)

	// ErrInternalServerError godoc.
	ErrInternalServerError = Failure(http.StatusInternalServerError, StatusInternalServerError)

	// ErrGatewayTimeout godoc.
	ErrGatewayTimeout = Failure(http.StatusGatewayTimeout, StatusGatewayTimeout)
)

// Err godoc.
func Err(err error) *JSON {
	return ErrInternalServerError.SetError(err)
}

// ErrByStatusCode godoc.
func ErrByStatusCode(statusCode int) *JSON {
	return ErrInternalServerError.SetError(errors.New(http.StatusText(statusCode)))
}

// Failure godoc.
func Failure(statusCode int, messageCode string) *JSON {
	res := JSON{
		Status:     FAILURE,
		Message:    StatusText(messageCode),
		StatusCode: statusCode,
	}

	res.Error = errors.New(res.Message.EN)
	return &res
}
