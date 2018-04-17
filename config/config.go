package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Port    uint   `env:"MEM_PORT" envDefault:"8080"`
	DataDir string `env:"MEM_DATADIR" envDefault:"memento_data/"`
}

func GetConfig() *Config {

	configuration := Config{}

	err := env.Parse(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return &configuration
}
