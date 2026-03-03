package di

import (
	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.OpenAIProvider,
	config.New,
)
