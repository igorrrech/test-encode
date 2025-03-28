package main

import (
	"context"
	"test/internal/config"
	"test/internal/http"
	"test/persondb"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.MustLoadConfig("./config.json")
	logger := logrus.New()

	sp := persondb.NewConnectionProvider(
		"postgres",
		cfg.Dsn,
		nil,
		logger,
	)

	pr := persondb.NewPersonRepository("persons")

	svc := http.NewService(
		cfg.Host,
		cfg.Port,
		logger,
		pr,
		sp,
	)

	ctx := context.Background()
	svc.Run(ctx)
}
