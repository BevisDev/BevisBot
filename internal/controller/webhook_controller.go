package controller

import (
	"net/http"

	"github.com/BevisDev/BevisBot/internal/service"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookController struct {
	botService service.BotService
}

func NewWebhookController(
	botService service.BotService,
) *WebhookController {
	return &WebhookController{
		botService: botService,
	}
}

func (w *WebhookController) Webhook(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
	go w.botService.HandleUpdate(update)
}
