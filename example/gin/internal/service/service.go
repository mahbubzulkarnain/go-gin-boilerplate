package service

import (
	"modul"
	"modul/example/gin/internal/repository"

	modulService "modul/service"
)

// V1 ...
type V1 struct {
	ModulService modul.Service
}

// Service ...
type Service struct {
	V1
}

// NewService ...
func NewService(r repository.Repository) Service {
	return Service{
		V1{
			ModulService: modulService.NewModulService(r.ModulRepository),
		},
	}
}
