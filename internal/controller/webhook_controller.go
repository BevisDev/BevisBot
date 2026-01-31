package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookController struct {
	started         map[int64]bool
	helloController *HelloController
}

func NewWebhookController(
	helloController *HelloController,
) *WebhookController {
	return &WebhookController{
		helloController: helloController,
	}
}

func (w *WebhookController) Webhook(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	go w.dispatch(update)
}

func (w *WebhookController) dispatch(update tgbotapi.Update) {
	msg := update.Message
	if msg == nil || !msg.IsCommand() {
		return
	}

	cmd := msg.Command()
	chatID := msg.Chat.ID
	//args := msg.CommandArguments()

	switch cmd {
	case "hello":
		w.helloController.Hello(chatID)
	}
}
