package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"

	"modul/example/gin/internal/delivery/http/middleware"
	"modul/example/gin/internal/service"
)

// Handler ...
type Handler struct {
	service service.Service
}

// RegisterServer ...
func RegisterServer(service service.Service) {
	gin.SetMode(envy.Get("GIN_MODE", gin.ReleaseMode))

	h := Handler{service: service}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Trace)

	r.POST("/", PlaceHolder)

	v1 := r.Group("/modul/v1")
	{
		v1.PATCH("/:id", h.UpdateByID)
		v1.DELETE("/:id", h.DeleteByID)
		v1.GET("/:id", h.FindOneByID)

		v1.GET("/", h.Find)
		v1.POST("/", h.Create)
	}

	PORT := ":" + envy.Get("PORT", "80")
	log.Println(fmt.Sprintf("0.0.0.0%v 🚀 server started... ", PORT))
	log.Fatal(r.Run(PORT))
}
