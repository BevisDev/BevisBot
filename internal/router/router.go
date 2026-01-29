package router

import (
	"github.com/BevisDev/BevisBot/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var APIPrefix = "/api"

func Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group(APIPrefix)
	{
		c := controller.NewWebhookController()
		api.POST("/webhook", c.Webhook)
	}
}
