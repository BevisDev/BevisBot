package notification

import (
	"github.com/BevisDev/BevisBot/internal/app/dto/request"
	"github.com/gin-gonic/gin"
)

type INotification interface {
	SendMessage(ctx gin.Context, r *request.NotificationRequest) error
}
