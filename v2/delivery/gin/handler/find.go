package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find ...
func (h Handler) Find(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
