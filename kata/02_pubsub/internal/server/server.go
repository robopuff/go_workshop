package server

import (
	"log"
	"net/http"
	"os"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/config"
)

func NewServer(cfg config.HTTP, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Bind,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		ErrorLog:     log.New(os.Stderr, "http: ", log.LstdFlags),
	}
}
