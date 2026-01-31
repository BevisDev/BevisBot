package middleware

import (
	"github.com/BevisDev/BevisBot/internal/controller"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookHandler struct {
	helloController *controller.HelloController
}

func NewWebhookHandler(
	helloController *controller.HelloController,
) *WebhookHandler {
	return &WebhookHandler{
		helloController: helloController,
	}
}

func (w *WebhookHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (w *WebhookHandler) dispatch(update tgbotapi.Update) {

}
