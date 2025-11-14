//go:build wireinject

package di

import (
	"github.com/BevisDev/BevisBot/internal/di/provider"
	"github.com/google/wire"
)

func NewServiceDI() *provider.ServiceProvider {
	wire.Build(
		NewNotificationServiceDI,
		provider.NewServiceProvider,
	)
	return new(provider.ServiceProvider)
}
