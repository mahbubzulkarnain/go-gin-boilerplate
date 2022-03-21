package response

import (
	"net/http"

	"github.com/gomodul/abstraction"
)

// Created ...
func Created(data interface{}) JSON {
	return JSON{
		Status:     SUCCESS,
		Message:    StatusText(StatusCreated),
		Data:       data,
		StatusCode: http.StatusCreated,
		Error:      nil,
	}
}

// Success ...
func Success(data interface{}) JSON {
	return SuccessWithPagination(data, nil)
}

// SuccessWithPagination ...
func SuccessWithPagination(data interface{}, info *abstraction.PaginationInfo) JSON {
	return JSON{
		Status:     SUCCESS,
		Message:    StatusText(StatusOK),
		Pagination: info,
		Data:       data,
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}
