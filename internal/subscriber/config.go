package subscriber

// Config configuration of subscriber.
type Config struct {
	Topics []string `koanf:"topics"`
	Count  int      `koanf:"count"`
}
