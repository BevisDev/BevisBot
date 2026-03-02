package main

import (
	"log"

	"github.com/BevisDev/BevisBot/internal/di"
)

func main() {
	app, err := di.InitializeApp()
	if err != nil {
		log.Fatalf("initialize app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run app %v", err)
	}
}
