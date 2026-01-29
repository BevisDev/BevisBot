package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	bot *tgbotapi.BotAPI
}

func New(token string) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	b := &TgBot{
		bot: bot,
	}

	log.Printf("bot started successfully")
	return b, nil
}
