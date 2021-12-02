package config

import "time"

type AMQP struct {
	Address string `env:"AMQP_ADDRESS" env-default:"localhost:5672"`
}

type HTTP struct {
	Bind         string        `env:"HTTP_BIND" env-default:":8080"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" env-default:"2s"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" env-default:"1s"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:"5s"`
}

type Config struct {
	Amqp AMQP
	Http HTTP
}
