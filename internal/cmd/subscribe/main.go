package subscribe

import (
	"github.com/mehditeymorian/qsse-test/internal/config"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/mehditeymorian/qsse-test/internal/subscriber"
	"github.com/spf13/cobra"
	"sync"
)

// Command subscribe command.
func Command() *cobra.Command {
	subscribe := &cobra.Command{
		Use:   "subscribe",
		Short: "subscribe to QSSE server",
		Run:   run,
	}

	return subscribe
}

func run(_ *cobra.Command, _ []string) {
	cfg := config.Load("config.yaml")

	log := logger.New(cfg.Logger)

	waitGroup := &sync.WaitGroup{}

	subscriber.Subscriber{
		WaitGroup:     waitGroup,
		Logger:        log,
		ServerAddress: cfg.Subscriber.ServerAddress,
	}.Subscribe(cfg.Subscriber)
}
