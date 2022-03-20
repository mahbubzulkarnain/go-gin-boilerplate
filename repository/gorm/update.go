package repository

import (
	"context"

	"modul"
	"modul/dto"
)

// Update ...
func (r *repository) Update(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	result := r.DB.WithContext(ctx).Save(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.ExecResponse{
		ID:           entity.ID,
		RowsAffected: result.RowsAffected,
	}, nil
}
