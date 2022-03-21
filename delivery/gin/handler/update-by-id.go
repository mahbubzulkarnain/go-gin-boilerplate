package handler

import (
	"github.com/gin-gonic/gin"

	"modul/dto"
	"modul/example/gin/pkg/response"
)

// UpdateByID ...
func (h Handler) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var payload dto.UpdateByIDRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.ErrBadRequest.JSON(c, payload)
		return
	}
	res, err := h.service.UpdateByID(c, id, payload)
	if err != nil {
		response.Err(err).JSON(c)
		return
	}
	response.Success(res).JSON(c)
}
