package repository

import "modul/v2"

type repository struct {
}

// NewModulRepository ...
func NewModulRepository() modul.Repository {
	return &repository{}
}
