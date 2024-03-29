package config

import (
	"errors"
	"flag"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Host                string
	Port                int
	BaseShortURLAddress string
}

func NewConfig() (Config, error) {
	config := Config{}

	serverAddress, ok := os.LookupEnv("SERVER_ADDRESS")
	if ok {
		host, port, err := parseServerAddress(serverAddress)
		if err != nil {
			return config, err
		}

		config.Host = host
		config.Port = port
	}

	flag.Func("a", "HTTP server address", func(address string) error {
		if config.Host != "" && config.Port != 0 {
			return nil
		}

		host, port, err := parseServerAddress(address)
		if err != nil {
			return err
		}

		config.Host = host
		config.Port = port

		return nil
	})

	baseURL, ok := os.LookupEnv("BASE_URL")
	if ok {
		_, err := url.ParseRequestURI(baseURL)
		if err != nil {
			return config, errors.New("need valid address for short URL in the format scheme://host:port/")
		}

		config.BaseShortURLAddress = baseURL
	}

	flag.Func("b", "Base address for short URL", func(flagValue string) error {
		if config.BaseShortURLAddress != "" {
			return nil
		}

		_, err := url.ParseRequestURI(flagValue)
		if err != nil {
			return errors.New("need valid address for short URL in the format scheme://host:port/")
		}

		config.BaseShortURLAddress = flagValue

		return nil
	})

	flag.Parse()

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == 0 {
		config.Port = 8080
	}

	if config.BaseShortURLAddress == "" {
		config.BaseShortURLAddress = "http://localhost:8080/"
	}

	return config, nil
}

func (c *Config) GetAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func parseServerAddress(address string) (host string, port int, err error) {
	splitAddress := strings.Split(address, ":")
	if len(splitAddress) != 2 {
		return "", 0, errors.New("need HTTP server address in the format host:port")
	}

	port, err = strconv.Atoi(splitAddress[1])
	if err != nil {
		return "", 0, err
	}

	return splitAddress[0], port, nil
}
