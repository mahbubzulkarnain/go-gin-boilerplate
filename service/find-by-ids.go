package service

import (
	"context"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// FindByIDs ...
func (s service) FindByIDs(ctx context.Context, ids []string) (data []*user.Entity, err error) {
	return s.userRepository.FindByIDs(ctx, ids)
}
