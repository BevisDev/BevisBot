package di

import (
	"github.com/BevisDev/BevisBot/internal/service"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	service.NewTaskService,
	service.NewBotService,
	service.NewOpenAIService,
)
