package handler

import (
	"github.com/gin-gonic/gin"

	"modul/example/gin/pkg/response"
)

// DeleteByID ...
func (h Handler) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.DeleteByID(c, id)
	if err != nil {
		response.Err(err).JSON(c)
		return
	}
	response.Success(res).JSON(c)
}
