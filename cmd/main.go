package main

import (
	"context"
	"test/internal/config"
	"test/internal/http"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.MustLoadConfig("./config.json")
	logger := logrus.New()

	svc := http.NewService(
		cfg.Host,
		cfg.Port,
		logger,
	)

	ctx := context.Background()
	svc.Run(ctx)
}
