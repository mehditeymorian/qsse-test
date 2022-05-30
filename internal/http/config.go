package http

// Config is the configuration for the HTTP server.
type Config struct {
	Port int `koanf:"port"`
}
