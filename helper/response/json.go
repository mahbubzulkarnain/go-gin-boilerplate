package response

import (
	"github.com/gin-gonic/gin"
)

// SUCCESS godoc.
const SUCCESS = "SUCCESS"

// PROGRESS godoc.
const PROGRESS = "PROGRESS"

// FAILURE godoc.
const FAILURE = "FAILURE"

// JSONMessage godoc.
type JSONMessage struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// JSON godoc.
type JSON struct {
	StatusCode int         `json:"-"`
	Error      error       `json:"-"`
	Status     string      `json:"status" example:"SUCCESS,PROGRESS,FAILURE"`
	Message    JSONMessage `json:"message"`
	Data       interface{} `json:"data"`
}

// SetError godo.
func (res JSON) SetError(err error) *JSON {
	res.Error = err
	return &res
}

// JSON godoc.
func (res JSON) JSON(c *gin.Context, data ...interface{}) {
	if res.Status == FAILURE && data != nil && data[0] != nil {
		res.Data = data[0]
	}

	c.AbortWithStatusJSON(res.StatusCode, res)
}
