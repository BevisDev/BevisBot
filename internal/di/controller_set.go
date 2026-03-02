package di

import (
	"github.com/BevisDev/BevisBot/internal/controller"
	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewWebhookController,
	controller.NewTaskController,
)
