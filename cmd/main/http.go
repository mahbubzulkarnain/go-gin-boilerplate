package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/delivery/http/handler"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/delivery/http/middleware"
)

func main() {
	gin.SetMode(envy.Get("GIN_MODE", gin.ReleaseMode))

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Trace)

	r.POST("/", handler.PlaceHolder)

	api := r.Group("/order")
	{
		v1 := api.Group("/v1")
		{
			v1.Any("/", handler.Ping)
		}
	}

	PORT := ":" + envy.Get("PORT", "80")
	log.Println(fmt.Sprintf("0.0.0.0%v 🚀 server started... ", PORT))
	log.Fatal(r.Run(PORT))
}
