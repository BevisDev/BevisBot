package controller

import (
	"net/http"

	"github.com/BevisDev/BevisBot/internal/bot"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookController struct{}

func NewWebhookController() *WebhookController {
	return &WebhookController{}
}

func (w *WebhookController) Webhook(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
	go bot.HandleUpdate(update)
}
