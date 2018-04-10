package config

import (
	"github.com/caarlos0/env"
	"log"
)

type Config struct {
	Port uint `env:"MEM_PORT" envDefault:"8080"`
}

func GetConfig() *Config {

	configuration := Config{}

	err := env.Parse(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return &configuration
}