package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/helper/response"
)

// Ping godoc.
func Ping(c *gin.Context) {
	response.Success(time.Now().Unix()).JSON(c)
}
