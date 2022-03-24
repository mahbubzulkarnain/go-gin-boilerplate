package repository

import (
	"gorm.io/gorm"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// UserModel ...
type UserModel struct {
	gorm.Model

	*user.Entity
}

// TableName ...
func (UserModel) TableName() string {
	return "users"
}

// BeforeCreate ...
func (u *UserModel) BeforeCreate(_ *gorm.DB) (err error) {
	return u.HashingPassword()
}
