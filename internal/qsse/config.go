package qsse

// Config configuration struct of QSSE.
type Config struct {
	Topics []string `koanf:"topics"`
	Port   int      `koanf:"port"`
}
