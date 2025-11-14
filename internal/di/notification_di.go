//go:build wireinject

package di

import (
	"github.com/BevisDev/BevisBot/internal/app/service/notification"
	"github.com/google/wire"
)

func NewNotificationServiceDI() notification.INotification {
	wire.Build(
		notification.NewNotification,
	)
	return new(notification.Notification)
}
