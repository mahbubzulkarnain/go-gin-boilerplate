package http

import (
	"github.com/gin-gonic/gin"

	"modul/example/gin/pkg/response"
)

// PlaceHolderPayload ...
type PlaceHolderPayload struct {
	StatusCode int `json:"status_code"`
}

// PlaceHolder ...
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
