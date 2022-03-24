package repository

import (
	"context"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// UpdateByID ...
func (r repository) UpdateByID(ctx context.Context, id string, e *user.Entity) (rowsAffected int64, err error) {
	result := r.DB.WithContext(ctx).Where("id=?", id).Updates(e)
	return result.RowsAffected, result.Error
}
