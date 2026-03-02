//go:build wireinject

package di

import (
	"github.com/BevisDev/BevisBot/internal"
	"github.com/BevisDev/BevisBot/internal/config"
	"github.com/BevisDev/BevisBot/internal/router"
	"github.com/google/wire"
)

func InitializeApp() (*internal.App, error) {
	wire.Build(
		controllerSet,
		serviceSet,
		repositorySet,
		config.New,
		router.New,
		internal.New,
	)
	return nil, nil
}
