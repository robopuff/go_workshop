package dictionary

import (
	"encoding/json"
	"net/http"

	"github.com/go-workshop/kata1/internal/server"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type handler struct {
	reader EntriesReader
}

func NewHandler(reader EntriesReader) server.Handler {
	return &handler{reader}
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	word, ok := vars["word"]
	if !ok || word == "" {
		w.WriteHeader(http.StatusBadRequest)
		server.JsonError(w, "word variable cannot be empty")
		return
	}

	entries, err := h.reader.Read(word)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		server.JsonError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(MapEntries(entries)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("Error while processing response JSON")
		server.JsonError(w, err.Error())
		return
	}
}
