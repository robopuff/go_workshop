package pool

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	ID string
	C  chan []byte

	p  Pool
	ws *websocket.Conn
}

func NewClient(id string, p Pool, ws *websocket.Conn) *Client {
	c := &Client{
		ID: id,
		C:  make(chan []byte),
		p:  p,
		ws: ws,
	}

	p.Register(c)
	return c
}

func (c *Client) WriterWorker() {
	// Ping time
	t := time.NewTicker(5 * time.Second)
	defer func() {
		t.Stop()
		c.ws.Close()
		c.p.Unregister(c)
	}()

	writeTimeout := 1 * time.Second
	for {
		select {
		case m, ok := <- c.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeTimeout))
			if !ok {
				if err := c.ws.WriteMessage(websocket.CloseMessage, nil); err != nil {
					logrus.WithError(err).Warn("close write message failed")
				}
				return
			}
			if err := c.ws.WriteMessage(websocket.TextMessage, m); err != nil {
				logrus.WithError(err).Warn("text write message failed")
				return
			}
		case <- t.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.WithError(err).Info("ping write message failed")
				return
			}
		}
	}
}
