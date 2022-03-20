package http

import (
	"github.com/gin-gonic/gin"

	"modul/dto"
	"modul/example/gin/pkg/response"
)

// Create ...
func (h Handler) Create(c *gin.Context) {
	var payload dto.CreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.ErrBadRequest.JSON(c, payload)
		return
	}
	res, err := h.service.ModulService.Create(c, payload)
	if err != nil {
		response.Err(err).JSON(c)
		return
	}
	response.Success(res).JSON(c)
}
