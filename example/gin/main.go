package main

import (
	"modul/example/gin/internal/delivery/http/handler"
	"modul/example/gin/internal/repository"
	"modul/example/gin/internal/service"
)

func main() {
	repo := repository.NewRepository(nil)
	s := service.NewService(repo)
	http.RegisterServer(s)
}
