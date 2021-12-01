package subscriber

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/server"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/subscriber/pool"
	"github.com/sirupsen/logrus"
)

type wsHandler struct {
	pool      pool.Pool
}

func NewWSHandler(pool pool.Pool) server.Handler {
	return &wsHandler{pool}
}

func (h *wsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(*http.Request) bool {
			// don't check origin - not safe but good enough for it right now
			return true
		},
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			logrus.WithError(reason).Errorf("websocket error %d", status)
		},
	}
	c, err := u.Upgrade(w, r, w.Header())
	if err != nil {
		logrus.WithError(err).Error("websocket upgrade errored")
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
		return
	}

	if err := c.WriteMessage(websocket.TextMessage, []byte("Welcome")); err != nil {
		logrus.WithError(err).Error("cannot write welcome message")
		c.Close()
		return
	}

	client := pool.NewClient("humpty bumpty", h.pool, c)
	client.WriterWorker()
}
