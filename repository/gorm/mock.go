package repository

import (
	"context"

	"github.com/gomodul/abstraction"
	"github.com/stretchr/testify/mock"

	"modul"
	"modul/dto"
)

// Mock ...
type Mock struct {
	mock.Mock
}

// Create ...
func (r *Mock) Create(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	args := r.Called(ctx, entity)
	return args.Get(0).(*dto.ExecResponse), args.Error(1)
}

// Update ...
func (r *Mock) Update(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	args := r.Called(ctx, entity)
	return args.Get(0).(*dto.ExecResponse), args.Error(1)
}

// Delete ...
func (r *Mock) Delete(ctx context.Context, entity *modul.Entity) (*dto.ExecResponse, error) {
	args := r.Called(ctx, entity)
	return args.Get(0).(*dto.ExecResponse), args.Error(1)
}

// Find ...
func (r *Mock) Find(ctx context.Context, filter *dto.Filter, p *abstraction.Pagination) ([]*modul.Entity, error) {
	args := r.Called(ctx, filter, p)
	return args.Get(0).([]*modul.Entity), args.Error(1)
}

// FindOne ...
func (r *Mock) FindOne(ctx context.Context, filter dto.Filter) (*modul.Entity, error) {
	args := r.Called(ctx, filter)
	return args.Get(0).(*modul.Entity), args.Error(1)
}
