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
	config := Config{
		Host:                "localhost",
		Port:                8080,
		BaseShortURLAddress: "http://localhost:8080/",
	}

	hostFlag := flag.String("host", "localhost", "HTTP server host")
	portFlag := flag.Int("port", 8080, "HTTP server port")
	baseURLFlag := flag.String("base-url", "http://localhost:8080/", "Base address for short URL")

	flag.Parse()

	config.Host = *hostFlag
	config.Port = *portFlag
	config.BaseShortURLAddress = *baseURLFlag

	return config
}

func (cfg *Config) GetAddress() string {
	return cfg.Host + ":" + strconv.Itoa(cfg.Port)
}

func (cfg *Config) GetBaseShortURLAddress() string {
	return cfg.BaseShortURLAddress
}
