package service

import (
	"context"

	"modul"
	"modul/dto"
)

// FindOneByID ...
func (s service) FindOneByID(ctx context.Context, id string) (*modul.Entity, error) {
	return s.modulRepository.FindOne(ctx, dto.Filter{
		ID: &id,
	})
}
