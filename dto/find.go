package dto

import (
	"github.com/gomodul/abstraction"
)

// FindRequest ...
type FindRequest struct {
	*Filter
	*abstraction.Pagination
}
