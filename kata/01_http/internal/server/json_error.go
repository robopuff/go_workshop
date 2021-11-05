package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JsonError(writer http.ResponseWriter, message interface{}) {
	e := json.NewEncoder(writer)
	if err := e.Encode(map[string]interface{}{
		"error": message,
	}); err != nil {
		logrus.WithError(err).Errorf("JSON error prasing")
		http.Error(writer, "error while processing JSON", http.StatusInternalServerError)
	}
}
