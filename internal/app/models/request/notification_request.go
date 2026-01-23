package request

import "github.com/BevisDev/BevisBot/internal/app/enums"

type NotificationRequest struct {
	Message  string
	NotiType enums.NotificationType
}
