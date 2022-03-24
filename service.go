package user

import (
	"context"

	"github.com/gomodul/abstraction"

	"github.com/mahbubzulkarnain/golang-boilerplate/user/dto"
)

// Service ...
type Service interface {
	Create(ctx context.Context, e *Entity) (rowsAffected int64, err error)
	UpdateByID(ctx context.Context, id string, e *Entity) (rowsAffected int64, err error)
	DeleteByIDs(ctx context.Context, ids []string) (rowsAffected int64, err error)

	Find(ctx context.Context, f *Filter, p *abstraction.Pagination) (data []string, info *abstraction.PaginationInfo, err error)
	FindByIDs(ctx context.Context, ids []string) (data []*Entity, err error)

	Login(ctx context.Context, args dto.LoginRequest, jwt dto.JWTConfig) (*dto.LoginResponse, error)
}
