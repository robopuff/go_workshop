package pool

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Pool interface {
	Register(*Client)
	Unregister(*Client)
	PropagateMessage([]byte)
}

type pool struct {
	clients sync.Map
}

func NewPool() Pool {
	return &pool{sync.Map{}}
}

func (p *pool) Register(client *Client) {
	logrus.Info("registered new client in pool")
	p.clients.Store(client, struct {}{})
}

func (p *pool) Unregister(client *Client) {
	logrus.Info("unregistered new client in pool")
	p.clients.Delete(client)
}

func (p *pool) PropagateMessage(m []byte) {
	p.clients.Range(func(k, _ interface{}) bool {
		client, ok := k.(*Client)
		if !ok {
			logrus.Warn("skipping client in pool, wrong type")
			return true
		}

		select {
		case client.C <- m:
			// Message have been propagated properly
		default:
			close(client.C)
			p.Unregister(client)
		}
		return true
	})
}


