package cron

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/BevisDev/BevisBot/internal/di/provider"
	"github.com/BevisDev/godev/utils"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron *cron.Cron
	deps *provider.ServiceProvider
}

func NewCron(
	deps *provider.ServiceProvider,
) *Cron {
	c := cron.New()
	return &Cron{
		cron: c,
		deps: deps,
	}
}

func (c *Cron) Start(ctx context.Context) {
	// register
	c.register()

	// start
	c.cron.Start()
	log.Println("Cron started success")

	go func() {
		<-ctx.Done()
		log.Println("cron is stopping")
		c.cron.Stop()
	}()
}

func (c *Cron) register() {
	for _, job := range c.GetJobs() {
		if job.disabled {
			log.Printf("cron %s is disabled\n", job.name)
			continue
		}

		_, err := c.cron.AddFunc(job.schedule,
			func() {
				ctx := utils.NewCtx()
				defer func() {
					if r := recover(); r != nil {
						trace := debug.Stack()
						log.Printf("panic recovered: %v, trace %s", r, trace)
					}
				}()

				job.handle(ctx)
			})
		if err != nil {
			log.Printf("error register %s cron: %v", job.name, err)
		}
	}
}
