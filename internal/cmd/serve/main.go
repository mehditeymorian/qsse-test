package serve

import (
	"strconv"

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

	server, err := qsse.NewServer(":"+strconv.Itoa(cfg.QSSE.Port), qsse.GetDefaultTLSConfig(), cfg.QSSE.Topics)
	if err != nil {
		log.Fatal("failed to run qsse server", zap.Error(err))
	}

	handler.Event{
		Server: server,
		Logger: log,
	}.Register(app)

	panic(app.Listen(":" + strconv.Itoa(cfg.HTTP.Port)))
}
