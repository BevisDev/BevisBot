package controller

import (
	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/view"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(chatID int64) {
	msg := view.Hello()
	lib.Bot.Send(chatID, msg)
}
