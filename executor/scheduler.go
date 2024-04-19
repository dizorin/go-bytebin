package executor

import (
	"context"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

var Scheduler *gocron.Scheduler

func SetupScheduler(ctx context.Context) {
	lc, _ := time.LoadLocation("UTC")
	Scheduler = gocron.NewScheduler(lc)

	period := utils.GetenvDuration("CHECK_PERIOD_FILE")
	_, err := Scheduler.Every(period).Do(checkInvalidationTask, ctx)
	if err != nil {
		log.Fatal(err)
	}

	Scheduler.StartAsync()
}

func checkInvalidationTask(ctx context.Context) {
	// TODO invalidate files
}
