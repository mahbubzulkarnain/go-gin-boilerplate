package response

import "net/http"

// Created godoc.
func Created(data interface{}) JSON {
	return JSON{
		Status:     SUCCESS,
		Message:    StatusText(StatusCreated),
		Data:       data,
		StatusCode: http.StatusCreated,
		Error:      nil,
	}
}

// Success godoc.
func Success(data interface{}) JSON {
	return JSON{
		Status:     SUCCESS,
		Message:    StatusText(StatusOK),
		Data:       data,
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}
