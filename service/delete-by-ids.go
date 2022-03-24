package service

import "context"

// DeleteByIDs ...
func (s service) DeleteByIDs(ctx context.Context, ids []string) (rowsAffected int64, err error) {
	return s.userRepository.DeleteByIDs(ctx, ids)
}
