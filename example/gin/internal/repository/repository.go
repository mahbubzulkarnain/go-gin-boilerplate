package repository

import (
	"gorm.io/gorm"
	"modul"

	modulRepository "modul/repository/gorm"
)

// V1 ...
type V1 struct {
	ModulRepository modul.Repository
}

// Repository ...
type Repository struct {
	V1
}

// NewRepository ...
func NewRepository(db *gorm.DB) Repository {
	return Repository{
		V1{
			ModulRepository: modulRepository.NewModulRepository(db),
		},
	}
}
