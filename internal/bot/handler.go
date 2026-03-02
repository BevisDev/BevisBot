package bot

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/lib/openai"
	"github.com/BevisDev/BevisBot/internal/view"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const sessionDuration = 1 * time.Hour

type Bot struct {
	mu            sync.RWMutex
	sessionExpiry map[int64]time.Time
}

func New() *Bot {
	return &Bot{
		sessionExpiry: make(map[int64]time.Time),
	}
}

func startSession(chatID int64) {
	sessionMu.Lock()
	defer sessionMu.Unlock()
	sessionExpiry[chatID] = time.Now().Add(sessionDuration)
}

func isSessionActive(chatID int64) bool {
	sessionMu.Lock()
	defer sessionMu.Unlock()
	exp, ok := sessionExpiry[chatID]
	if !ok {
		return false
	}
	if time.Now().After(exp) {
		delete(sessionExpiry, chatID)
		return false
	}
	return true
}

func HandleUpdate(update tgbotapi.Update) {
	if lib.Bot == nil {
		return
	}
	msg := update.Message
	if msg == nil {
		return
	}
	chatID := msg.Chat.ID

	if msg.IsCommand() {
		cmd := msg.Command()
		log.Printf("[bot] chat=%d cmd=/%s", chatID, cmd)
		handleCommand(chatID, cmd, msg.CommandArguments())
		return
	}

	// Tin nhắn thường: chỉ trả lời khi đang trong session (sau /hello, trong 1h)
	if !isSessionActive(chatID) {
		return
	}
	userText := msg.Text
	if userText == "" {
		return
	}
	replyWithOpenAI(chatID, userText)
}

func handleCommand(chatID int64, cmd, args string) {
	switch cmd {
	case "start", "hello":
		startSession(chatID)
		lib.Bot.Send(chatID, view.Hello())
	case "help":
		lib.Bot.Send(chatID, view.Help())
	default:
		lib.Bot.Send(chatID, "Lệnh chưa được hỗ trợ. Gõ /help để xem hướng dẫn.")
	}
	_ = args
}

func replyWithOpenAI(chatID int64, userMessage string) {
	if lib.OpenAIAPIKey == "" {
		lib.Bot.Send(chatID, "Chưa cấu hình OpenAI. Admin cần set openai.apiKey trong config.")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()
	reply, err := openai.Ask(ctx, lib.OpenAIAPIKey, userMessage)
	if err != nil {
		log.Printf("[bot] openai error: %v", err)
		lib.Bot.Send(chatID, "Đang bận, bạn thử lại sau nhé.")
		return
	}
	// Telegram giới hạn 4096 ký tự / tin
	if len(reply) > 4000 {
		reply = reply[:4000] + "..."
	}
	lib.Bot.Send(chatID, reply)
}
