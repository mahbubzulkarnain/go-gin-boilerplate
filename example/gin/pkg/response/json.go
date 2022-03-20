package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodul/abstraction"
)

// SUCCESS ...
const SUCCESS = "SUCCESS"

// PROGRESS ...
const PROGRESS = "PROGRESS"

// FAILURE ...
const FAILURE = "FAILURE"

// JSONMessage ...
type JSONMessage struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// JSON ...
type JSON struct {
	StatusCode int                         `json:"-"`
	Error      error                       `json:"-"`
	Status     string                      `json:"status" example:"SUCCESS,PROGRESS,FAILURE"`
	Pagination *abstraction.PaginationInfo `json:"pagination,omitempty"`
	Message    JSONMessage                 `json:"message"`
	Data       interface{}                 `json:"data"`
}

// SetError godo.
func (res JSON) SetError(err error) *JSON {
	res.Error = err
	return &res
}

// JSON ...
func (res JSON) JSON(c *gin.Context, data ...interface{}) {
	if res.Status == FAILURE && data != nil && data[0] != nil {
		res.Data = data[0]
	}

	c.AbortWithStatusJSON(res.StatusCode, res)
}
