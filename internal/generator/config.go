package generator

// Config is a struct that holds the configuration for the generator service.
type Config struct {
	Generators          []EventGenerator `koanf:"generators"`
	IdenticalGenerators bool             `koanf:"identical_generators"`
	Count               int              `koanf:"count"`
	PublishAddress      string           `koanf:"publish_address"`
	Timeout             int              `koanf:"timeout"`
}

// EventGenerator is a struct that holds the configuration for a single generator.
type EventGenerator struct {
	Topic     string `koanf:"topic"`
	EventSize int    `koanf:"event_size"`
	Rate      int    `koanf:"rate"`
}
