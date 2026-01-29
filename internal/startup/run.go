package startup

import (
	"context"
	"fmt"
	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/lib/tgbot"
	"github.com/BevisDev/BevisBot/internal/router"
	"github.com/BevisDev/godev/framework"
	"github.com/BevisDev/godev/ginfw/server"
	"github.com/BevisDev/godev/logger"
	"github.com/BevisDev/godev/rest"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Run() {
	if err := config.Load(); err != nil {
		log.Fatalf("error load config failed %v", err)
	}

	isProd := strings.HasPrefix(config.AppConfig.Server.Profile, "prod")
	isDev := strings.HasPrefix(config.AppConfig.Server.Profile, "dev")

	options := []framework.Option{
		framework.WithLogger(&logger.Config{
			IsProduction: isProd,
			IsLocal:      isDev,
			DirName:      config.AppConfig.Logger.DirName,
			Filename:     config.AppConfig.Logger.FileName,
		}),
		framework.WithRESTClient(
			rest.WithSkipHeader(),
		),
		framework.WithServer(&server.Config{
			IsProduction: isProd,
			Port:         "8080",
			Setup: func(r *gin.Engine) {
				router.Register(r)
			},
		}),
	}
	b := framework.New(options...)
	b.AfterInit(func(ctx context.Context) error {
		bot, err := tgbot.New(config.AppConfig.TgBot.Token)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
		lib.Bot = bot
		return nil
	})
	b.Run(context.Background())
}
