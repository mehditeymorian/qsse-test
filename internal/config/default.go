package config

import (
	"github.com/mehditeymorian/qsse-test/internal/generator"
	"github.com/mehditeymorian/qsse-test/internal/http"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/mehditeymorian/qsse-test/internal/qsse"
	"github.com/mehditeymorian/qsse-test/internal/subscriber"
)

// Default returns the default configurations.
func Default() Config {
	return Config{
		HTTP: http.Config{
			Port: 8080,
		},
		QSSE: qsse.Config{
			Topics: []string{"topic"},
			Port:   4242,
		},
		Logger: logger.Config{
			Level: "debug",
		},
		Generator: generator.Config{
			Generators: []generator.EventGenerator{
				{
					Topic:     "topic",
					EventSize: 100,
					Rate:      500,
				},
			},
			IdenticalGenerators: true,
			Count:               1,
			PublishAddress:      "http://localhost:8080/event/publish",
			Timeout:             5000,
		},
		Subscriber: subscriber.Config{
			Topics:        []string{"topic"},
			Count:         1,
			ServerAddress: "localhost:4242",
		},
	}
}
