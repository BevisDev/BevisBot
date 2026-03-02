package internal

import (
	"context"
	"fmt"
	"strings"

	bot2 "github.com/BevisDev/BevisBot/internal/bot"
	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/lib/tgbot"
	"github.com/BevisDev/BevisBot/internal/router"
	"github.com/BevisDev/godev/framework"
	"github.com/BevisDev/godev/ginfw/server"
	"github.com/BevisDev/godev/logger"
	"github.com/BevisDev/godev/rest"
	"github.com/BevisDev/godev/utils"
	"github.com/gin-gonic/gin"
)

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *config.Config
	router *router.Router
}

func New(
	config *config.Config,
	router *router.Router,
) (*App, error) {
	ctx, cancel := utils.NewCtxCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
		config: config,
		router: router,
	}, nil
}

func (a *App) Run() error {
	cfg := a.config.AppConfig()

	isProd := strings.HasPrefix(cfg.Server.Profile, "prod")
	isDev := strings.HasPrefix(cfg.Server.Profile, "dev")

	options := []framework.Option{
		framework.WithLogger(&logger.Config{
			IsProduction: isProd,
			IsLocal:      isDev,
			DirName:      cfg.Logger.DirName,
			Filename:     cfg.Logger.FileName,
		}),
		framework.WithRESTClient(
			rest.WithSkipHeader(),
		),
		framework.WithServer(&server.Config{
			IsProduction: isProd,
			Port:         "8080",
			Setup: func(r *gin.Engine) {
				a.router.Register(r)
			},
		}),
	}
	b := framework.New(a.ctx, options...)

	//b.AddServices(func(ctx context.Context) error {
	//	dbCfg := cfg.Database
	//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
	//		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.DB, dbCfg.Port)
	//
	//	db, err := db.New(dsn)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	lib.DB = db
	//	return nil
	//})

	b.AddServices(func(ctx context.Context) error {
		bot, err := tgbot.New(cfg.TgBot.Token)
		if err != nil {
			return fmt.Errorf("tg bot init: %w", err)
		}
		lib.Bot = bot
		lib.OpenAIAPIKey = cfg.OpenAI.APIKey
		return nil
	})

	b.AfterInit(func(ctx context.Context) error {
		if isDev {
			go func() {
				lib.Bot.StartLongPolling(ctx, bot2.HandleUpdate)
			}()
		}
		return nil
	})

	if err := b.Run(a.ctx); err != nil {
		return err
	}
	return nil
}
