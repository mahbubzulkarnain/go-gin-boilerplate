package handler

import (
	"github.com/gin-gonic/gin"

	"modul/example/gin/pkg/response"
)

// FindOneByID ...
func (h Handler) FindOneByID(c *gin.Context) {
	id := c.Param("id")

	entity, err := h.service.FindOneByID(c, id)
	if err != nil {
		response.Err(err).JSON(c)
		return
	}
	response.Success(entity).JSON(c)
}
