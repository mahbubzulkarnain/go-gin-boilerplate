package repository

import (
	"context"

	"modul"
	"modul/dto"
)

// Create ...
func (r *repository) Create(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	result := r.DB.WithContext(ctx).Create(&ModulModel{Entity: entity})
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.ExecResponse{
		ID:           entity.ID,
		RowsAffected: result.RowsAffected,
	}, nil
}
