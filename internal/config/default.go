package config

import (
	"github.com/mehditeymorian/qsse-test/internal/http"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/mehditeymorian/qsse-test/internal/qsse"
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
	}
}
