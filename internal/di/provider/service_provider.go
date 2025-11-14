package provider

import "github.com/BevisDev/BevisBot/internal/app/service/notification"

type ServiceProvider struct {
	NotiService notification.INotification
}

func NewServiceProvider(
	NotiService notification.INotification,
) *ServiceProvider {
	return &ServiceProvider{
		NotiService: NotiService,
	}
}
