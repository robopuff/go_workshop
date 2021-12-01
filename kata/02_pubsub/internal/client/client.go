package client

type Closer interface {
	Close()
}

type Consumer interface {
	Read() (string, error)
}

type Publisher interface {
	Send(message string) error
}

type ConsumerPublisher interface {
	Consumer
	Publisher
	Closer
}
