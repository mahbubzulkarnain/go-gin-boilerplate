package repository

import (
	"context"

	"modul"
	"modul/dto"
)

// FindOne ...
func (r repository) FindOne(ctx context.Context, filter dto.Filter) (data *modul.Entity, err error) {
	err = r.DB.WithContext(ctx).Where(filter).Take(&data).Error
	return
}
