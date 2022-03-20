package service

import (
	"context"

	"modul/dto"
)

// UpdateByID ...
func (s service) UpdateByID(ctx context.Context, id string, args dto.UpdateByIDRequest) (*dto.ExecResponse, error) {
	entity, err := s.modulRepository.FindOne(ctx, dto.Filter{ID: &id})
	if err != nil {
		return nil, err
	}
	entity.Name = args.Name
	return s.modulRepository.Update(ctx, entity)
}
