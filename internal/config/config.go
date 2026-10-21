package config

import (
	"os"
)

type Config struct {
	Token     string
	ServiceID string
}

func Load() Config {
	return Config{
		Token:     os.Getenv("TOKEN"),
		ServiceID: os.Getenv("SERVICE_ID"),
	}
}
