package publisher

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type wsHandler struct {
	publisher rabbitmq.Publisher
}

func NewWSHandler(publisher rabbitmq.Publisher) server.Handler {
	return &wsHandler{publisher}
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

	defer c.Close()
	for {
		mt, m, err := c.ReadMessage()
		if err != nil {
			logrus.WithError(err).Error("websocket read error")
			break
		}

		if err := h.publisher.Publish(
			m,
			[]string{"queue:chat"},
			rabbitmq.WithPublishOptionsContentType("plain/text"),
		); err != nil {
			logrus.WithError(err).Error("cannot publish message in AMQP")
		}

		err = c.WriteMessage(mt, []byte(fmt.Sprintf("produced message with `%s`", m)))
		if err != nil {
			logrus.WithError(err).Error("websocket write error")
			break
		}
	}
}
