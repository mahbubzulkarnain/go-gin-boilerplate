package service

import (
	"context"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// UpdateByID ...
func (s service) UpdateByID(ctx context.Context, id string, e *user.Entity) (int64, error) {
	return s.userRepository.UpdateByID(ctx, id, e)
}
