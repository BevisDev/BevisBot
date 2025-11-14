package cron

import (
	"context"
	"log"

	"github.com/BevisDev/BevisBot/internal/app/service/notification"
)

type ReportDailyJob struct {
	notiService notification.INotification
}

func NewReportDailyJob(
	notiService notification.INotification,
) IJob {
	return &ReportDailyJob{
		notiService: notiService,
	}
}

func ReportDailyCron(ctx context.Context) {
	log.Printf("Job ReportDailyCron is running...")
}

func (r *ReportDailyJob) Run(ctx context.Context) {
	log.Printf("Job ReportDailyCron is running...")
}
