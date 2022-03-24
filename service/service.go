package service

import (
	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

type service struct {
	userRepository user.Repository
}

// NewUserService ...
func NewUserService(userRepository user.Repository) user.Service {
	return &service{
		userRepository: userRepository,
	}
}
