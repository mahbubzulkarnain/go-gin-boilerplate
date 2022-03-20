package repository

import (
	"context"

	"github.com/gomodul/abstraction"

	"modul"
	"modul/dto"
)

// Find ...
func (r repository) Find(ctx context.Context, filter *dto.Filter, p *abstraction.Pagination) (data []*modul.Entity, err error) {
	err = r.DB.WithContext(ctx).Scopes(Pagination(p)).Where(filter).Find(&data).Error
	return
}
