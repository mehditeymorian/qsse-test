package logger

type Config struct {
	Level string `validate:"required" koanf:"level"`
}
