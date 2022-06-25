package subscribe

import (
	"net/http"
	"sync"

	"github.com/mehditeymorian/qsse-test/internal/config"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/mehditeymorian/qsse-test/internal/subscriber"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
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

	metric := subscriber.NewMetrics("client")

	waitGroup := &sync.WaitGroup{}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()

	subscriber.Subscriber{
		WaitGroup:     waitGroup,
		Logger:        log,
		ServerAddress: cfg.Subscriber.ServerAddress,
		Metric:        metric,
	}.Subscribe(cfg.Subscriber)
}
