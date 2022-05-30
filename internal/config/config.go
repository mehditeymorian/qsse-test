package config

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/mehditeymorian/qsse-test/internal/generator"
	"github.com/mehditeymorian/qsse-test/internal/http"
	"github.com/mehditeymorian/qsse-test/internal/logger"
	"github.com/mehditeymorian/qsse-test/internal/qsse"
	"github.com/mehditeymorian/qsse-test/internal/subscriber"
	"github.com/tidwall/pretty"
)

// Prefix env variable prefix.
const Prefix = "QSSETEST"

// Config contains all the configuration for the application.
type Config struct {
	HTTP       http.Config       `koanf:"http"`
	Logger     logger.Config     `koanf:"logger"`
	QSSE       qsse.Config       `koanf:"qsse"`
	Generator  generator.Config  `koanf:"generator"`
	Subscriber subscriber.Config `koanf:"subscriber"`
}

// Load loads the config from default, the given path, and env variables.
func Load(path string) Config {
	var cfg Config

	koan := koanf.New(".")

	// load default configuration
	if err := koan.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	// load configuration from file
	if err := koan.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml: %v", err)
	}

	// load environment variables
	callback := func(key string, value string) (string, interface{}) {
		finalKey := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(key, Prefix)), "__", ".")

		if strings.Contains(value, ",") {
			// remove all the whitespace from value
			// split the value using comma
			finalValue := strings.Split(removeWhitespace(value), ",")

			return finalKey, finalValue
		}

		return finalKey, value
	}
	if err := koan.Load(env.ProviderWithValue(Prefix, ".", callback), nil); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	if err := koan.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	indent, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatalf("error marshal config: %v", err)
	}

	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(cfgStrTemplate, string(indent))

	return cfg
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}
