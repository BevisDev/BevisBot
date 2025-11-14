package startup

import (
	"context"

	"github.com/BevisDev/BevisBot/internal/cron"
	"github.com/BevisDev/BevisBot/internal/di"
)

func Initialize(ctx context.Context) error {
	// init cron
	deps := di.NewServiceDI()
	c := cron.NewCron(deps)
	c.Start(ctx)

	return nil
}
