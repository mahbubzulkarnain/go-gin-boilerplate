package repository

import (
	"context"
	"modul/dto"

	"modul"
)

// Delete ...
func (r *repository) Delete(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	result := r.DB.WithContext(ctx).Delete(&ModulModel{Entity: entity})
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.ExecResponse{
		ID:           entity.ID,
		RowsAffected: result.RowsAffected,
	}, nil
}
