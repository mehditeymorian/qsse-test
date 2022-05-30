package generator

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/mehditeymorian/qsse-test/internal/http/request"
	"go.uber.org/zap"
)

// Generator struct for creating event generators.
type Generator struct {
	Logger         *zap.Logger
	WaitGroup      *sync.WaitGroup
	PublishAddress string
	Client         *http.Client
}

// Generate creates event generators.
func (g Generator) Generate(cfg Config) {
	if cfg.IdenticalGenerators {
		if len(cfg.Generators) < 1 {
			g.Logger.Fatal("No generators specified")
		}

		for i := 0; i < cfg.Count; i++ {
			go g.generateEvents(cfg.Generators[0].Topic, cfg.Generators[0].Rate, cfg.Generators[0].EventSize)
		}
		g.Logger.Info("created generators", zap.Int("count", cfg.Count), zap.String("topic", cfg.Generators[0].Topic))
	} else {
		for _, generator := range cfg.Generators {
			go g.generateEvents(generator.Topic, generator.Rate, generator.EventSize)
			g.Logger.Info("created generator", zap.String("topic", generator.Topic))
		}
	}

	g.WaitGroup.Wait()
}

func (g Generator) generateEvents(topic string, rate, eventSize int) {
	g.WaitGroup.Add(1)

	for {
		req := request.Publish{
			Topic:   topic,
			Message: randomText(eventSize),
		}

		data, err := json.Marshal(req)
		if err != nil {
			g.Logger.Warn("Failed to marshal request", zap.Error(err))

			continue
		}

		reader := bytes.NewReader(data)
		resp, err := g.Client.Post(g.PublishAddress, "application/json", reader)

		switch {
		case err != nil:
			g.Logger.Warn("Failed to send publish request", zap.String("topic", topic), zap.Error(err))
		case resp.StatusCode >= 400:
			g.Logger.Warn("request unsuccessful", zap.String("topic", topic), zap.Int("status", resp.StatusCode))
		default:
			g.Logger.Info("event sent to publish", zap.String("topic", topic))
		}

		<-time.Tick(time.Duration(rate) * time.Millisecond) // nolint:staticcheck
	}
}

func randomText(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	data := make([]byte, length)

	for i := range data {
		data[i] = characters[rand.Intn(len(characters))] //nolint:gosec
	}

	return string(data)
}
