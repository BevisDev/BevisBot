package router

import (
	"net/http"

	"github.com/BevisDev/BevisBot/internal/controller"
	"github.com/gin-gonic/gin"
)

var APIPrefix = "/api"

func Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	helloController := controller.NewHelloController()
	c := controller.NewWebhookController(
		helloController,
	)

	api := r.Group(APIPrefix)
	{
		api.POST("/webhook", c.Webhook)
	}
}
