package server

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HealthData struct {
	Host string `json:"host"`
	Time time.Time `json:"server_time"`
}

type healthz struct{}

func NewHealthzHandler() Handler {
	return &healthz{}
}

func (h healthz) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	host, _ := os.Hostname()
	json.NewEncoder(w).Encode(HealthData{
		Host: host,
		Time: time.Now(),
	})
}
