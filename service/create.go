package service

import (
	"context"

	"modul"
	"modul/dto"
)

// Create ...
func (s service) Create(ctx context.Context, args dto.CreateRequest) (*dto.ExecResponse, error) {
	return s.modulRepository.Create(ctx, &modul.Entity{
		Name: args.Name,
	})
}
