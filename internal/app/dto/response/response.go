package response

import (
	"github.com/BevisDev/BevisBot/internal/app/enums"
	"github.com/BevisDev/BevisBot/internal/helper"
	"github.com/BevisDev/godev/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  true,
		Data:       data,
		ResponseAt: helper.GetNowDateTime(),
	})
}

func Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  true,
		ResponseAt: helper.GetNowDateTime(),
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  true,
		Data:       data,
		ResponseAt: helper.GetNowDateTime(),
	})
}

func Unauthorized(c *gin.Context, code enums.ResponseCode) {
	c.JSON(http.StatusUnauthorized, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  false,
		ResponseAt: helper.GetNowDateTime(),
		Error: &Error{
			ErrorCode: code,
			Message:   code.Message(),
		},
	})
}

func BadRequest(c *gin.Context, code enums.ResponseCode, message string) {
	c.JSON(http.StatusBadRequest, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  false,
		ResponseAt: helper.GetNowDateTime(),
		Error: &Error{
			ErrorCode: code,
			Message:   message,
		},
	})
}

func ServerError(c *gin.Context, code enums.ResponseCode) {
	c.JSON(http.StatusInternalServerError, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  false,
		ResponseAt: helper.GetNowDateTime(),
		Error: &Error{
			ErrorCode: code,
			Message:   code.Message(),
		},
	})
}

func SetErrorCode(c *gin.Context, httpCode int, code enums.ResponseCode) {
	c.JSON(httpCode, &Response{
		State:      utils.GetState(c.Request.Context()),
		IsSuccess:  false,
		ResponseAt: helper.GetNowDateTime(),
		Error: &Error{
			ErrorCode: code,
			Message:   code.Message(),
		},
	})
}
