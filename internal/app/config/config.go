package config

import (
	"flag"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Address             string
	BaseShortURLAddress string
}

func NewConfig() (Config, error) {
	config := Config{
		Address:             "localhost:8080",
		BaseShortURLAddress: "http://localhost:8080/",
	}

	flag.StringVar(&config.Address, "address", "localhost:8080", "HTTP server address")
	flag.StringVar(&config.BaseShortURLAddress, "base-url", "http://localhost:8080/", "Base address for short URL")

	flag.Parse()

	if address, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		config.Address = address
	}

	if baseURL, ok := os.LookupEnv("BASE_URL"); ok {
		config.BaseShortURLAddress = baseURL
	}

	err := config.validate()
	if err != nil {
		return Config{}, err
	}

	config.normalize()

	return config, nil
}

func (c *Config) validate() error {
	_, err := http.NewRequest("GET", c.Address, nil)
	if err != nil {
		return err
	}

	_, err = http.NewRequest("GET", c.BaseShortURLAddress, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) normalize() {
	c.BaseShortURLAddress = strings.TrimSuffix(c.BaseShortURLAddress, "/")
}
