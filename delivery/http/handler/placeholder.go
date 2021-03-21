package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/helper/response"
)

// PlaceHolderPayload godoc.
type PlaceHolderPayload struct {
	StatusCode int `json:"status_code"`
}

// PlaceHolder godoc.
func PlaceHolder(c *gin.Context) {
	var payload PlaceHolderPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.ErrBadRequest.JSON(c, payload)
		return
	}

	response.JSON{
		Status:     response.SUCCESS,
		Data:       payload,
		StatusCode: payload.StatusCode,
	}.JSON(c)
}
