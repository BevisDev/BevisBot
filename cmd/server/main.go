package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BevisDev/godev/consts"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	r := gin.Default()

	var s = consts.Email
	fmt.Print(s)

	_, err := tgbotapi.NewBotAPI(s)
	if err != nil {
		log.Panic(err)
	}

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
