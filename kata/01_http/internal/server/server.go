package server

import (
	"net/http"
	"time"
)

func NewServer(handler http.Handler, addr string) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}
