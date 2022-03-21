package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"

	"example/gin/pkg/middleware"

	moduldelivery "github.com/username/modul/delivery/gin/handler"
	modulrepository "github.com/username/modul/repository/gorm"
	modulservice "github.com/username/modul/service"

	moduldelivery2 "github.com/username/modul/v2/delivery/gin/handler"
	modulrepository2 "github.com/username/modul/v2/repository/kafka"
	modulservice2 "github.com/username/modul/v2/service"
)

func main() {
	gin.SetMode(envy.Get("GIN_MODE", gin.ReleaseMode))

	modulRepository := modulrepository.NewModulRepository(nil)
	modulRepository2 := modulrepository2.NewModulRepository()

	modulService := modulservice.NewModulService(modulRepository)
	modulService2 := modulservice2.NewModulService(modulRepository2)

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Trace)

	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	moduldelivery.RegisterServer(r, modulService)
	moduldelivery2.RegisterServer(r, modulService2)

	PORT := ":" + envy.Get("PORT", "80")
	log.Println(fmt.Sprintf("0.0.0.0%v 🚀 server started... ", PORT))
	log.Fatal(r.Run(PORT))
}
