package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// singleton design pattern
var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	DB          *DBConfig
}

func mustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		fmt.Printf("❌ missing required env var: %s\n", key)
		os.Exit(1)
	}
	return val
}

func mustGetInt(key string) int {
	valStr := mustGet(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		fmt.Printf("❌ failed to convert %s to int\n", key)
		os.Exit(1)
	}
	return val
}

func mustGetBool(key string) bool {
	valStr := mustGet(key)
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		fmt.Printf("❌ failed to convert %s to bool\n", key)
		os.Exit(1)
	}
	return val
}

func loadConfig() {
	// .env load
	_ = godotenv.Load()

	dbConfig := &DBConfig{
		Host:          mustGet("DB_HOST"),
		Port:          mustGetInt("DB_PORT"),
		Name:          mustGet("DB_NAME"),
		User:          mustGet("DB_USER"),
		Password:      mustGet("DB_PASSWORD"),
		EnableSSLMode: mustGetBool("DB_ENABLE_SSL_MODE"),
	}

	configurations = &Config{
		Version:     mustGet("VERSION"),
		ServiceName: mustGet("SERVICE_NAME"),
		HttpPort:    mustGetInt("HTTP_PORT"),
		DB:          dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
