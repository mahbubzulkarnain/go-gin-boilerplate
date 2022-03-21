package handler

import (
	"github.com/gin-gonic/gin"

	"modul/v2"
)

// Handler ...
type Handler struct {
	service modul.Service
}

// RegisterServer ...
func RegisterServer(r *gin.Engine, service modul.Service) {
	h := Handler{service: service}

	api := r.Group("/modul/v2")
	{
		api.GET("/", h.Find)
	}
}
