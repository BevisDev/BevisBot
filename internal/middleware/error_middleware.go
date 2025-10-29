package middleware

import (
	"context"
	"fmt"
	"github.com/BevisDev/BevisBot/internal/app/dto/response"
	"github.com/BevisDev/BevisBot/internal/app/enums"
	"github.com/BevisDev/godev/consts"
	"github.com/BevisDev/godev/utils/random"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// store state in context
		var state = random.RandUUID()
		ctx := context.WithValue(c.Request.Context(), consts.State, state)
		c.Request = c.Request.WithContext(ctx)

		defer func() {
			if err := recover(); err != nil {
				trace := debug.Stack()
				fmt.Printf("panic recovered: %v, trace: %s", err, trace)
				//lib.Logger.Error(state, "Panic occurred: {}, trace: {}", err, trace)
				response.ServerError(c, enums.ServerError)
				c.Abort()
				return
			}
		}()

		c.Next()

		if len(c.Errors) != 0 {
			err := c.Errors.Last().Err
			response.BadRequest(c, enums.InvalidRequest, err.Error())
			return
		}
	}
}
