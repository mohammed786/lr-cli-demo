package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env"
)

// Config represents env vars required for app.
// Incase no env is supplied, default values are used as fallback.
type Config struct {
	LoginRadiusAPIKey    string `env:"LoginRadiusAPIKey" envDefault:"ddff8a63-cbc3-4723-8415-b910c4d8770d"`
	LoginRadiusAPIDomain string `env:"LoginRadiusAPIDomain" envDefault:"https://devapi.lrinternal.com"`
	AdminConsoleAPIDomain string `env:"AdminConsoleAPIDomain" envDefault:"https://devadmin-console-api.lrinternal.com"`
}

// Singleton instance of config
var instance *Config

var once sync.Once

// Read and parse the configuration file
func read() *Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

// GetInstance returns the singleton instance of `Config`
func GetInstance() *Config {
	once.Do(func() {
		instance = read()
	})
	return instance
}
