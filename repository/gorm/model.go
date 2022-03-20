package repository

import (
	"gorm.io/gorm"

	"modul"
)

// ModulModel ...
type ModulModel struct {
	gorm.DB

	*modul.Entity
}

// TableName ...
func (ModulModel) TableName() string {
	return "modul"
}
