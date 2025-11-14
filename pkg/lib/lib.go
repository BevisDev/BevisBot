package lib

import (
	"DecisionEngine/src/main/infrastructure/config"
	"context"

	"github.com/BevisDev/godev/logger"
)

var (
	Logger logger.Exec
)

func InitLogger(ctx context.Context) {
	cf := config.SystemConfig

	Logger = logger.New(&logger.Config{
		Profile:    cf.ServerConfig.Profile,
		DirName:    cf.LoggerConfig.DirName,
		Filename:   cf.LoggerConfig.FileName,
		MaxSize:    cf.LoggerConfig.MaxSize,
		MaxBackups: cf.LoggerConfig.MaxBackups,
		MaxAge:     cf.LoggerConfig.MaxAge,
		Compress:   cf.LoggerConfig.Compress,
		IsSplit:    cf.LoggerConfig.IsSplit,
		CallerConfig: logger.CallerConfig{
			Request: logger.SkipGroup{
				Internal: 1,
				External: 4,
			},
			Response: logger.SkipGroup{
				Internal: 1,
				External: 5,
			},
		},
	})
	Logger.Info(state, "Logger starts success {}", true)
}
