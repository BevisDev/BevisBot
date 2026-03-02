package lib

import (
	"github.com/BevisDev/BevisBot/internal/lib/tgbot"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Bot          *tgbot.TgBot
	OpenAIAPIKey string
)
