package generate

import (
	"net/http"
	"sync"
	"time"

	"github.com/mehditeymorian/qsse-test/internal/config"
	"github.com/mehditeymorian/qsse-test/internal/generator"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/spf13/cobra"
)

// Command returns sub-command for serving server.
func Command() *cobra.Command {
	generate := &cobra.Command{ //nolint:exhaustruct
		Use:   "generate",
		Short: "generate events and publish to server",
		Long:  "generate events with different sizes with different rates to publish to QSSE server",
		Run:   run,
	}

	return generate
}

func run(_ *cobra.Command, _ []string) {
	cfg := config.Load("config.yaml")

	log := logger.New(cfg.Logger)

	group := &sync.WaitGroup{}

	generator.Generator{
		Logger:         log,
		WaitGroup:      group,
		PublishAddress: cfg.Generator.PublishAddress,
		Client: &http.Client{ //nolint:exhaustruct
			Timeout: time.Duration(cfg.Generator.Timeout) * time.Millisecond,
		},
	}.Generate(cfg.Generator)
}
