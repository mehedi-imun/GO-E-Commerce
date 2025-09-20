package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


// singleton design pattern
var configurations *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("faied to load env")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("faied to load version")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("faied to load serviceName")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("faied to load httpPort")
		os.Exit(1)
	}
	httpPortInt, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("failed to convert httpPort to int")
		os.Exit(1)
	}
	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPortInt,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
