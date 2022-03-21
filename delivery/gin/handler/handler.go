package handler

import (
	"github.com/gin-gonic/gin"

	"modul"
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
		api.PATCH("/:id", h.UpdateByID)
		api.DELETE("/:id", h.DeleteByID)
		api.GET("/:id", h.FindOneByID)

		api.GET("/", h.Find)
		api.POST("/", h.Create)
	}
}
