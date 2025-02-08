package app

import (
	"context"

	"go.uber.org/zap"
)

type App struct {
	logger  *zap.SugaredLogger
	sender  sender
	scraper scraper
}

type sender interface {
	RunSender(context.Context) error
	Stop() error
}

type scraper interface {
	RunScraper(context.Context) error
	Stop() error
}

func NewApp(logger *zap.SugaredLogger, sender sender, scraper scraper) *App {
	return &App{logger, sender, scraper}
}

func (a *App) Run(ctx context.Context) error {

	return nil
}
