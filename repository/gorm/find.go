package repository

import (
	"context"

	"github.com/gomodul/abstraction"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// Find ...
func (r repository) Find(ctx context.Context, f *user.Filter, p *abstraction.Pagination) (data []string, err error) {
	err = r.DB.WithContext(ctx).Scopes(Pagination(p)).Where(f).Find(&data).Error
	return
}
