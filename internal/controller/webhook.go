package controller

import (
	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookController struct {
}

func NewWebhookController() *WebhookController {
	return &WebhookController{}
}

func (c *WebhookController) Webhook(ctx *gin.Context) {
	var update tgbotapi.Update
	_ = ctx.ShouldBindJSON(&update)

	ctx.JSON(200, gin.H{"ok": true})

	if update.Message == nil {
		return
	}

	go handleMessage(update.Message)
}

var started = map[int64]bool{}

func handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	if msg.IsCommand() && msg.Command() == "start" {
		started[chatID] = true

		lib.Bot.Send(chatID, "Chào bạn 👋 Bot đã sẵn sàng!")
		return
	}
}
