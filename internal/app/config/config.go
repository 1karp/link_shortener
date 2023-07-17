package config

import (
	"flag"
	"strconv"
)

type Config struct {
	Host                string
	Port                int
	BaseShortURLAddress string
}

func NewConfig() Config {
	config := Config{}

	flag.StringVar(&config.Host, "host", "localhost", "HTTP server host")
	flag.IntVar(&config.Port, "port", 8080, "HTTP server port")
	flag.StringVar(&config.BaseShortURLAddress, "base-url", "", "Base address for short URL")

	flag.Parse()

	return config
}

func (c *Config) GetAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func (c *Config) GetBaseShortURLAddress() string {
	return c.BaseShortURLAddress
}
