package repository

import (
	"gorm.io/gorm"

	"modul"
)

type repository struct {
	DB *gorm.DB
}

// NewModulRepository ...
func NewModulRepository(db *gorm.DB) modul.Repository {
	if err := db.AutoMigrate(&ModulModel{}); err != nil {
		panic(err)
	}
	return &repository{
		DB: db,
	}
}
