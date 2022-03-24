package service

import (
	"context"

	"github.com/gomodul/abstraction"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// Find ...
func (s service) Find(ctx context.Context, f *user.Filter, p *abstraction.Pagination) ([]string, *abstraction.PaginationInfo, error) {
	entities, err := s.userRepository.Find(ctx, f, p)
	if err != nil {
		return nil, nil, err
	}
	if p != nil {
		data, info := p.NewPaginationInfo(entities)
		return data.([]string), info, nil
	}
	return entities, nil, nil
}
