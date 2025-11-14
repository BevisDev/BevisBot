package notification

import (
	"github.com/BevisDev/BevisBot/internal/app/dto/request"
	"github.com/gin-gonic/gin"
)

type Notification struct {
}

func NewNotification() INotification {
	return &Notification{}
}

func (n Notification) SendMessage(ctx gin.Context, r *request.NotificationRequest) error {
	//TODO implement me
	panic("implement me")
}
