package service

import (
	"context"

	"github.com/gomodul/abstraction"

	"modul"
	"modul/dto"
)

// Find ...
func (s service) Find(ctx context.Context, filter *dto.Filter, p *abstraction.Pagination) ([]*modul.Entity, *abstraction.PaginationInfo, error) {
	entities, err := s.modulRepository.Find(ctx, filter, p)
	if err != nil {
		return nil, nil, err
	}
	if p != nil {
		data, info := p.NewPaginationInfo(entities)
		return data.([]*modul.Entity), info, nil
	}
	return entities, nil, nil
}
