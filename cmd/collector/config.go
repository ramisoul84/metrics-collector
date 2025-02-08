package main

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel          string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	EnableDebugServer bool   `env:"ENABLE_DEBUG_SERVER" envDefault:"true"`
	DebugServerAddr   string `env:"DEBUG_SERVER_ADDR" envDefault:"0.0.0.0:8181"`
}

func ReadConfig() error {
	config := Config{}
	env.Parse(&config)

	return nil
}
