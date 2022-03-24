package service

import (
	"context"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// Create ...
func (s service) Create(ctx context.Context, e *user.Entity) (rowsAffected int64, err error) {
	return s.userRepository.Create(ctx, e)
}
