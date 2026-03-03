package service

import (
	"context"
	"log"

	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/view"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotService interface {
	HandleUpdate(update tgbotapi.Update)
}

type botService struct {
	openAIService OpenAIService
}

func NewBotService(
	openAIService OpenAIService,
) BotService {
	return &botService{
		openAIService: openAIService,
	}
}

func (b *botService) HandleUpdate(update tgbotapi.Update) {
	msg := update.Message
	if msg == nil {
		return
	}
	chatID := msg.Chat.ID
	chatType := msg.Chat.Type

	if msg.IsCommand() {
		cmd := msg.Command()
		log.Printf("[bot] chat=%d type=%s cmd=/%s", chatID, chatType, cmd)
		b.handleCommand(chatID, cmd, msg.CommandArguments())
		return
	}
}

func (b *botService) handleCommand(chatID int64, cmd, text string) {
	switch cmd {
	case "ask":
		reply, err := b.openAIService.Reply(context.Background(), text)
		if err != nil {
			lib.Bot.Send(chatID, "Đang bận, bạn thử lại sau nhé.")
		}

		if len(reply) > 4000 {
			reply = reply[:4000] + "..."
		}
		lib.Bot.Send(chatID, reply)
	case "help":
		lib.Bot.Send(chatID, view.Help())
	default:
		lib.Bot.Send(chatID, "Lệnh chưa được hỗ trợ. Gõ /help để xem hướng dẫn.")
	}
}
