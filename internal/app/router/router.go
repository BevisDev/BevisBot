package router

import (
	"github.com/BevisDev/BevisBot/internal/app/dto/response"
	"github.com/BevisDev/BevisBot/internal/app/enums"
	"github.com/gin-gonic/gin"
	"net/http"
)

const API = "/api"

func RegisterRouter(r *gin.Engine) {
	// public router

	// Handle undefined routes
	r.NoRoute(func(c *gin.Context) {
		//state := utils.GetState(c.Request.Context())
		//lib.Logger.Warn(state, "Route not found: %s %s", c.Request.Method, c.Request.URL.Path)
		response.SetErrorCode(c, http.StatusNotFound, enums.NotFound)
	})

	// Handle undefined HTTP methods
	r.NoMethod(func(c *gin.Context) {
		//state := utils.GetState(c.Request.Context())
		//lib.Logger.Warn(state, "Method not allowed: %s %s", c.Request.Method, c.Request.URL.Path)
		response.SetErrorCode(c, http.StatusMethodNotAllowed, enums.InvalidRequest)
	})
}
