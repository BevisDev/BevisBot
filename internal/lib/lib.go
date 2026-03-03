package lib

import (
	"github.com/BevisDev/godev/rest"
	"github.com/BevisDev/godev/tgbot"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Bot        *tgbot.TgBot
	RESTClient *rest.Client
)
