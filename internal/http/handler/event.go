package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	return ctx.SendStatus(http.StatusOK)
}
