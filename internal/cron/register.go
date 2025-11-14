package cron

import (
	"context"

	"github.com/BevisDev/BevisBot/internal/app/config"
)

type IJob interface {
	Run(ctx context.Context)
}

type Job struct {
	name     string
	schedule string
	disabled bool
	handle   func(ctx context.Context)
}

func (c *Cron) GetJobs() []*Job {
	var cf = config.AppConfig.Cron

	return []*Job{
		{
			name:     "ReportDaily",
			schedule: cf.ReportDaily.Cron,
			disabled: cf.IsDisabled,
			handle:   c.GetHandler("reportdaily"),
		},
	}
}

func (c *Cron) GetHandler(name string) func(ctx context.Context) {
	switch name {
	case "ReportDaily":
		return NewReportDailyJob(c.deps.NotiService).Run
	default:
		return nil
	}
}
