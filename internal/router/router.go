package router

import (
	"net/http"

	"github.com/BevisDev/BevisBot/internal/controller"
	"github.com/gin-gonic/gin"
)

var APIPrefix = "/api"

type Router struct {
	webhookCtrl *controller.WebhookController
	taskCtrl    *controller.TaskController
}

func New(
	webhookCtrl *controller.WebhookController,
	taskCtrl *controller.TaskController,
) *Router {
	return &Router{
		webhookCtrl: webhookCtrl,
		taskCtrl:    taskCtrl,
	}
}

func (rt *Router) Register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group(APIPrefix)
	{
		webhook(api, rt.webhookCtrl)
		task(api, rt.taskCtrl)
	}
}

func webhook(g *gin.RouterGroup, c *controller.WebhookController) {
	w := g.Group("/webhook")
	{
		w.POST("/", c.Webhook)
	}
}

func task(g *gin.RouterGroup, c *controller.TaskController) {
	//w := g.Group("/tasks")
}
