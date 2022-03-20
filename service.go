package modul

import (
	"context"

	"github.com/gomodul/abstraction"

	"modul/dto"
)

// Service ...
type Service interface {
	Create(ctx context.Context, args dto.CreateRequest) (*dto.ExecResponse, error)
	UpdateByID(ctx context.Context, id string, args dto.UpdateByIDRequest) (*dto.ExecResponse, error)
	DeleteByID(ctx context.Context, id string) (*dto.ExecResponse, error)

	Find(ctx context.Context, filter *dto.Filter, p *abstraction.Pagination) ([]*Entity, *abstraction.PaginationInfo, error)
	FindOneByID(ctx context.Context, id string) (*Entity, error)
}
