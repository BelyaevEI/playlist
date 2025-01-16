package main

import (
	"context"
	"log"

	"github.com/BelyaevEI/playlist/internal/app"
	"github.com/BelyaevEI/playlist/internal/logger"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	logger.Info("shutdown server")
}
