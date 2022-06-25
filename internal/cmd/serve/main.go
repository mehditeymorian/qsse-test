package serve

import (
	"strconv"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/qsse-test/internal/config"
	"github.com/mehditeymorian/qsse-test/internal/http/handler"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/snapp-incubator/qsse"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Command returns sub-command for serving server.
func Command() *cobra.Command {
	serve := &cobra.Command{ //nolint:exhaustruct
		Use:   "serve",
		Short: "server for publishing events into QSSE",
		Long:  "server for publishing events into QSSE",
		Run:   run,
	}

	return serve
}

func run(_ *cobra.Command, _ []string) {
	cfg := config.Load("config.yaml")

	log := logger.New(cfg.Logger)

	app := fiber.New()

	serverConfig := &qsse.ServerConfig{
		Metric: &qsse.MetricConfig{
			NameSpace: "qsset",
		},
		TLSConfig: qsse.GetDefaultTLSConfig(),
		Worker: &qsse.WorkerConfig{
			CleaningInterval:          time.Duration(cfg.QSSE.CleaningInterval) * time.Second,
			ClientAcceptorCount:       int64(cfg.QSSE.ClientAcceptorCount),
			ClientAcceptorQueueSize:   cfg.QSSE.ClientAcceptorQueueSize,
			EventDistributorCount:     int64(cfg.QSSE.EventDistributorCount),
			EventDistributorQueueSize: cfg.QSSE.EventDistributorQueueSize,
		},
	}

	server, err := qsse.NewServer(":"+strconv.Itoa(cfg.QSSE.Port), cfg.QSSE.Topics, serverConfig)
	if err != nil {
		log.Fatal("failed to run qsse server", zap.Error(err))
	}

	handler.Event{
		Server: server,
		Logger: log,
	}.Register(app)

	app.Get("/metrics", adaptor.HTTPHandler(server.MetricHandler()))

	panic(app.Listen(":" + strconv.Itoa(cfg.HTTP.Port)))
}
