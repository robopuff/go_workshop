package config

import "time"

type AMQP struct {
	Address string `env:"AMQP_ADDRESS"`
}

type HTTP struct {
	Bind         string        `env:"HTTP_BIND"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT"`
}

type Config struct {
	Amqp AMQP
	Http HTTP
}
