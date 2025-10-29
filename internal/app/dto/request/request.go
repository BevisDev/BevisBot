package request

import (
	"github.com/BevisDev/BevisBot/internal/app/consts"
	"github.com/gin-gonic/gin"
)

type Request[T any] struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	Page      int     `json:"page,omitempty"`
	Size      int     `json:"size,omitempty"`
	Search    string  `json:"search,omitempty"`
	Data      T       `json:"data,omitempty"`
}

func (r *Request[T]) GetBody(c *gin.Context) (*T, error) {
	if err := c.BindJSON(r.Data); err != nil {
		c.Error(err)
		return nil, err
	}
	return &r.Data, nil
}

func (r *Request[T]) GetCreatedBy() *string {
	if r.CreatedBy == nil {
		c := consts.SYSTEM
		return &c
	}
	return r.CreatedBy
}

func (r *Request[T]) GetUpdatedBy() *string {
	if r.UpdatedBy == nil {
		c := consts.SYSTEM
		return &c
	}
	return r.UpdatedBy
}
