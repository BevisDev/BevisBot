package internal

import (
	"context"
	"strings"

	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/router"
	"github.com/BevisDev/BevisBot/internal/service"
	"github.com/BevisDev/godev/framework"
	"github.com/BevisDev/godev/ginfw/server"
	"github.com/BevisDev/godev/logger"
	"github.com/BevisDev/godev/rest"
	"github.com/BevisDev/godev/tgbot"
	"github.com/BevisDev/godev/utils"
	"github.com/gin-gonic/gin"
)

type App struct {
	ctx        context.Context
	cancel     context.CancelFunc
	config     *config.Config
	router     *router.Router
	botService service.BotService
}

func New(
	config *config.Config,
	router *router.Router,
	botService service.BotService,
) (*App, error) {
	ctx, cancel := utils.NewCtxCancel(context.Background())
	return &App{
		ctx:        ctx,
		cancel:     cancel,
		config:     config,
		router:     router,
		botService: botService,
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
		framework.WithTgBot(&tgbot.Config{
			Token: cfg.TgBot.Token,
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

	b.AfterInit(func(ctx context.Context) error {
		lib.Bot = b.TgBot()
		lib.RESTClient = b.RESTClient()
		if isDev {
			go func() {
				lib.Bot.StartLongPolling(ctx,
					a.botService.HandleUpdate,
				)
			}()
		}
		return nil
	})

	if err := b.Run(a.ctx); err != nil {
		return err
	}
	return nil
}
