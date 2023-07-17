package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Host                string
	Port                int
	BaseShortURLAddress string
}

func NewConfig() Config {
	hostFlag := flag.String("a", "localhost", "HTTP server host")
	portFlag := flag.Int("p", 8080, "HTTP server port")
	baseURLFlag := flag.String("b", "http://localhost:8080/", "Base address for short URL")

	flag.Parse()

	return Config{
		Host:                *hostFlag,
		Port:                *portFlag,
		BaseShortURLAddress: *baseURLFlag,
	}
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) GetBaseShortURLAddress() string {
	return c.BaseShortURLAddress
}
