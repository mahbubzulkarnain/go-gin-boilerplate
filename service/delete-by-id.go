package service

import (
	"context"

	"modul/dto"
)

// DeleteByID ...
func (s service) DeleteByID(ctx context.Context, id string) (*dto.ExecResponse, error) {
	entity, err := s.modulRepository.FindOne(ctx, dto.Filter{ID: &id})
	if err != nil {
		return nil, err
	}
	return s.modulRepository.Delete(ctx, entity)
}
