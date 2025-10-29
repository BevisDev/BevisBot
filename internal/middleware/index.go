package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddleware(r gin.IRoutes) {
	r.Use(
		loggerHandler(),
		authHandler(),
		//serverTimeoutHandler(),
		errorHandler(),
	)
}
