package service

import "modul/v2"

type service struct {
	modulRepository modul.Repository
}

// NewModulService ...
func NewModulService(modulRepository modul.Repository) modul.Service {
	return &service{
		modulRepository: modulRepository,
	}
}
