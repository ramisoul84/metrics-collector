package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel           string        `env:"LOG_LEVEL" envDefault:"DEBUG"`
	EnableDebugServer  bool          `env:"ENABLE_DEBUG_SERVER" envDefault:"true"`
	DebugServerAddr    string        `env:"DEBUG_SERVER_ADDR" envDefault:"0.0.0.0:8181"`
	TargetHost         string        `env:"TARGET_HOST" envDefault:"0.0.0.0:8181"`
	UserToken          string        `env:"USER_TOKEN" envDefault:"0.0.0.0:8181"`
	SendMetricsTimeout time.Duration `env:"SEND_METRICS_TIMEOUT" envDefault:"15s"`
	ScrapeInterval     time.Duration `env:"SCRAPE_INTERVAL" envDefault:"1s"`
	ScrapeConfigPath   string        `env:"SCRAPE_CONFIG_PATH" envDefault:"0.0.0.0:8181"`
}

func ReadConfig() (*Config, error) {
	config := Config{}
	//err := os.Setenv("TARGET_HOST", "/tmp/fakehome")
	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("read config err: %w", err)
	}

	// Check required fields
	if config.TargetHost == "" {
		return nil, errors.New("TargetHost is required")
	}
	if config.UserToken == "" {
		return nil, errors.New("UserToken is required")
	}
	if config.ScrapeConfigPath == "" {
		return nil, errors.New("ScrapeConfigPath is required")
	}

	return &config, nil
}
