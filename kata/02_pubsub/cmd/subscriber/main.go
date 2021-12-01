package main

import (
	"os"
	"time"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/config"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/subscriber"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{
		Amqp: config.AMQP{
			Address: "localhost:5672",
		},
		Http: config.HTTP{
			Bind:         ":8090",
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
	}

	app := subscriber.NewApp(cfg.Amqp, cfg.Http)
	if err := app.Run(); err != nil {
		logrus.WithError(err).Error("error while running application")
	}

	os.Exit(app.ExitCode())
}
