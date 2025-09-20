package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPortStr := os.Getenv("HTTP_PORT")

	if version == "" || serviceName == "" || httpPortStr == "" {
		return nil, fmt.Errorf("missing required env variables")
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_PORT: %w", err)
	}

	return &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}, nil
}
