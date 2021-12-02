package subscriber

import (
	"fmt"
	"net/http"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/config"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/server"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/subscriber/pool"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/wagslane/go-rabbitmq"
)

const (
	ExitCodeOk = iota
	ExitCodeHTTPErr = 100
	ExitCodeAMQPErr = 200
)

type app struct {
	amqp     config.AMQP
	http     config.HTTP
	exitCode int
}

func NewApp(amqp config.AMQP, http config.HTTP) internal.App {
	return &app{amqp, http, ExitCodeOk}
}

func (a *app) Run() error {
	logrus.Info("starting ...")

	client, err := rabbitmq.NewConsumer(fmt.Sprintf("amqp://%s", a.amqp.Address), amqp.Config{})
	if err != nil {
		a.exitCode = ExitCodeAMQPErr
		return err
	}

	logrus.Info("AMQP connected")

	p := pool.NewPool()
	if err := client.StartConsuming(func(d rabbitmq.Delivery) bool {
		p.PropagateMessage(d.Body)
		return true
	}, "queue:chat", []string{}); err != nil {
		a.exitCode = ExitCodeAMQPErr
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", server.NewHealthHandler("subscriber", "dev").Handle)
	mux.HandleFunc("/ws", NewWSHandler(p).Handle)

	srv := server.NewServer(a.http, mux)
	logrus.WithField("bind", srv.Addr).Info("HTTP server starting")
	if err := srv.ListenAndServe(); err != nil {
		a.exitCode = ExitCodeHTTPErr
		return err
	}
	return nil
}

func (a app) ExitCode() int {
	return a.exitCode
}
