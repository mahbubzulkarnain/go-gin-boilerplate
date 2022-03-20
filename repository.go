package modul

import (
	"context"

	"github.com/gomodul/abstraction"

	"modul/dto"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, entity *Entity) (*dto.ExecResponse, error)
	Update(ctx context.Context, entity *Entity) (*dto.ExecResponse, error)
	Delete(ctx context.Context, entity *Entity) (*dto.ExecResponse, error)

	Find(ctx context.Context, filter *dto.Filter, p *abstraction.Pagination) ([]*Entity, error)
	FindOne(ctx context.Context, filter dto.Filter) (*Entity, error)
}
