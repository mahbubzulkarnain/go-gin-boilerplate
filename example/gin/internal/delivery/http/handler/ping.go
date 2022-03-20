package http

import (
	"time"

	"github.com/gin-gonic/gin"

	"modul/example/gin/pkg/response"
)

// Ping ...
func Ping(c *gin.Context) {
	response.Success(time.Now().Unix()).JSON(c)
}
