package client

import "github.com/gorilla/websocket"

type websockets struct {
	url  string
	conn *websocket.Conn
}

func NewWebsocketsClient(url string) ConsumerPublisher {
	if url == "" {
		panic("no url provided for publisher")
	}
	return &websockets{
		url: url,
	}
}

func (w *websockets) connection() (*websocket.Conn, error) {
	if w.conn != nil {
		return w.conn, nil
	}

	conn, _, err := websocket.DefaultDialer.Dial(w.url, nil)
	if err != nil {
		return nil, err
	}
	w.conn = conn
	return w.conn, nil
}

func (w *websockets) Close() {
	conn, err := w.connection()
	if err != nil {
		w.conn = nil
		return
	}

	conn.Close()
	w.conn = nil
}

func (w *websockets) Read() (string, error) {
	conn, err := w.connection()
	if err != nil {
		return "", err
	}

	_, message, err := conn.ReadMessage()
	return string(message), err
}

func (w *websockets) Send(message string) error {
	conn, err := w.connection()
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}
