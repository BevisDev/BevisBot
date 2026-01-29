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

func (t *TgBot) Send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	if _, err := t.bot.Send(msg); err != nil {
		log.Println("send error:", err)
	}
}
