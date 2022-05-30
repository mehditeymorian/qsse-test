package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/qsse-test/internal/http/request"
	"github.com/snapp-incubator/qsse"
	"go.uber.org/zap"
)

// Event is the handler for the event endpoints.
type Event struct {
	Server qsse.Server
	Logger *zap.Logger
}

// Register registers the event handler routes.
func (e Event) Register(app *fiber.App) {
	app.Post("/event/publish", e.publish)
}

func (e Event) publish(ctx *fiber.Ctx) error {
	e.Logger.Info("Received event publish request")

	var req request.Publish
	if err := ctx.BodyParser(&req); err != nil {
		e.Logger.Warn("failed to parse request body", zap.Error(err))

		return ctx.Status(http.StatusBadRequest).JSON(err) //nolint:wrapcheck
	}

	e.Logger.Info("publishing event", zap.String("topic", req.Topic))

	e.Server.Publish(req.Topic, []byte(req.Message))

	return ctx.SendStatus(http.StatusOK) //nolint:wrapcheck
}
