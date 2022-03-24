package user

import (
	"context"

	"github.com/gomodul/abstraction"

	"github.com/mahbubzulkarnain/golang-boilerplate/user/dto"
)

// Repository ...
type Repository interface {
	KeyByID(id string) string
	Create(ctx context.Context, e *Entity) (rowsAffected int64, err error)
	UpdateByID(ctx context.Context, id string, e *Entity) (rowsAffected int64, err error)
	DeleteByIDs(ctx context.Context, ids []string) (rowsAffected int64, err error)

	Find(ctx context.Context, f *Filter, p *abstraction.Pagination) (data []string, err error)
	FindByIDs(ctx context.Context, ids []string) ([]*Entity, error)

	Login(ctx context.Context, args dto.LoginRequest) (data *Entity, err error)
}
