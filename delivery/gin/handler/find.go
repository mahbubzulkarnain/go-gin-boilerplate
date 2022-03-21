package handler

import (
	"github.com/gin-gonic/gin"

	"modul/dto"
	"modul/example/gin/pkg/response"
)

// Find ...
func (h Handler) Find(c *gin.Context) {
	var request dto.FindRequest
	if err := c.BindQuery(&request); err != nil {
		response.Err(err).JSON(c)
		return
	}
	entities, info, err := h.service.Find(c, request.Filter, request.Pagination)
	if err != nil {
		response.Err(err).JSON(c)
		return
	}
	response.SuccessWithPagination(entities, info).JSON(c)
}
