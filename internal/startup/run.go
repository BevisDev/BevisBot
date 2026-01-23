package startup

import (
	"context"
	"strings"

	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/BevisDev/godev/framework"
	"github.com/BevisDev/godev/logx"
	"github.com/BevisDev/godev/rest"
)

// Run starts the application, sets up signal handling, and ensures graceful shutdown.
func Run() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		return
	}

	prod := strings.HasPrefix(appConfig.Server.Profile, "prod")
	dev := strings.HasPrefix(appConfig.Server.Profile, "dev")
	fw := framework.New(
		framework.WithLogger(&logx.Config{
			IsProduction: prod,
			IsLocal:      dev,
			MaxSize:      appConfig.Logger.MaxSize,
			MaxBackups:   appConfig.Logger.MaxBackups,
			MaxAge:       appConfig.Logger.MaxAge,
			Compress:     true,
			IsRotate:     true,
			DirName:      appConfig.Logger.DirName,
			Filename:     appConfig.Logger.FileName,
		}),
		framework.WithRestClient(
			rest.WithSkipHeader(),
		),
	)

	if err := fw.Run(context.Background()); err != nil {
		return
	}

}
