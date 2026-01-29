package response

import (
	"github.com/BevisDev/BevisBot/internal/app/enums"
)

type Response struct {
	State      string      `json:"state,omitempty" example:"8137ce10-305b-42f5-8f14-9c48dd6f23f0"`
	IsSuccess  bool        `json:"isSuccess" example:"true"`
	Data       interface{} `json:"data,omitempty"`
	ResponseAt string      `json:"responseAt,omitempty" example:"2025-01-14 16:44:47.510"`
	Error      *Error      `json:"error,omitempty"`
}

type Error struct {
	ErrorCode enums.ResponseCode `json:"code,omitempty" example:"3000"`
	Message   string             `json:"message,omitempty" example:"Invalid Request"`
}
