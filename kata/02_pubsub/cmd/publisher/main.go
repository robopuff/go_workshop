package main

import (
	"os"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/config"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/publisher"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		logrus.WithError(err).Error("cannot read env")
		os.Exit(-1000)
	}

	app := publisher.NewApp(cfg.Amqp, cfg.Http)
	if err := app.Run(); err != nil {
		logrus.WithError(err).Error("error while running application")
	}

	os.Exit(app.ExitCode())
}
