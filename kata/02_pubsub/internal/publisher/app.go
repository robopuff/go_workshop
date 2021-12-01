package publisher

import (
	"fmt"
	"net/http"
	"os"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/config"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wagslane/go-rabbitmq"
)

type app struct {
	amqp     config.AMQP
	http     config.HTTP
	exitCode int
}

func NewApp(amqp config.AMQP, http config.HTTP) internal.App {
	return &app{amqp, http, 0}
}

func (a *app) Run() error {
	logrus.Info("starting ...")

	client, returns, err := rabbitmq.NewPublisher(fmt.Sprintf("amqp://%s", a.amqp.Address), amqp.Config{})
	if err != nil {
		logrus.WithError(err).Error("cannot start AMQP client")
		os.Exit(-1)
	}

	logrus.Info("AMQP connected")

	go func() {
		for r := range returns {
			logrus.WithField("return", r).Info("AMQP server return message")
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", server.NewHealthHandler("publisher", "dev").Handle)
	mux.HandleFunc("/ws", NewWSHandler(client).Handle)

	srv := server.NewServer(a.http, mux)
	logrus.WithField("bind", srv.Addr).Info("HTTP server starting")
	if err := srv.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("HTTP server critical error")
		os.Exit(500)
	}
	return nil
}

func (a app) ExitCode() int {
	return a.exitCode
}
