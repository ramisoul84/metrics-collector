package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel           string        `env:"LOG_LEVEL" envDefault:"DEBUG"`
	EnableDebugServer  bool          `env:"ENABLE_DEBUG_SERVER" envDefault:"true"`
	DebugServerAddr    string        `env:"DEBUG_SERVER_ADDR" envDefault:"0.0.0.0:8181"`
	TargetHost         string        `env:"TARGET_HOST" envDefault:"host-domain.com:9999"`
	UserToken          string        `env:"USER_TOKEN" `
	ScrapeInterval     time.Duration `env:"SCRAPE_INTERVAL" envDefault:"15s"`
	SendMetricsTimeout time.Duration `env:"SEND_METRICS_TIMEOUT"`
	ScrapeConfigpath   string        `env:"SCRAPE_CONFIG_PATH"`
}

func ReadConfig() (*Config, error) {
	config := Config{}
	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("error reading config %w", err)
	}

	return &config, nil
}
