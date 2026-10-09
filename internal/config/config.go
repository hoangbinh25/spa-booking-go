package config

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort int

	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	appPort, err := getPort("APP_PORT", "8080")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := getPort("DB_PORT", "5432")
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: appPort,

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBName:     getEnv("DB_NAME", "spa_booking"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) DatabaseURL() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.DBUser, c.DBPassword),
		Host:     net.JoinHostPort(c.DBHost, strconv.Itoa(c.DBPort)),
		Path:     c.DBName,
		RawQuery: url.Values{"sslmode": {c.DBSSLMode}}.Encode(),
	}
	return u.String()
}

func (c Config) Validate() error {
	var validationErrors []error

	require := func(key, value string) {
		if strings.TrimSpace(value) == "" {
			validationErrors = append(validationErrors, fmt.Errorf("%s must be set", key))
		}
	}

	require("DB_HOST", c.DBHost)
	require("DB_NAME", c.DBName)
	require("DB_USER", c.DBUser)
	require("DB_PASSWORD", c.DBPassword)

	if c.AppPort < 1 || c.AppPort > 65535 {
		validationErrors = append(validationErrors, fmt.Errorf("APP_PORT must be between 1 and 65535"))
	}
	if c.DBPort < 1 || c.DBPort > 65535 {
		validationErrors = append(validationErrors, fmt.Errorf("DB_PORT must be between 1 and 65535"))
	}

	validSSLModes := map[string]bool{
		"disable": true, "allow": true, "prefer": true, "require": true,
		"verify-ca": true, "verify-full": true,
	}
	if !validSSLModes[c.DBSSLMode] {
		validationErrors = append(validationErrors, fmt.Errorf("DB_SSLMODE must be one of disable, allow, prefer, require, verify-ca, verify-full"))
	}

	return errors.Join(validationErrors...)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getPort(key, fallback string) (int, error) {
	value := getEnv(key, fallback)
	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid number, got %q", key, value)
	}
	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("%s must be between 1 and 65535, got %d", key, port)
	}
	return port, nil
}
