package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type response struct {
	Name       string
	ServerTime time.Time
	Version    string
}

type health struct {
	name, version string
}

func NewHealthHandler(name, version string) Handler {
	return &health{name, version}
}

func (h health) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response{
		Name:       h.name,
		ServerTime: time.Now(),
		Version:    h.version,
	}); err != nil {
		logrus.WithError(err).Error("cannot marshal JSON response")
		http.Error(w, "cannot marshal response", http.StatusInternalServerError)
	}
}
