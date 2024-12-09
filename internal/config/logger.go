package config

import (
	"errors"
	"os"
)

const (
	logLevelEnvName = "LOGGER"
)

type LoggerConfig interface {
	Level() string
}

type loggerConfig struct {
	logLevel string
}

func parseLogLevelenv() (*loggerConfig, error) {
	logLevel := os.Getenv(logLevelEnvName)
	if len(logLevel) == 0 {
		return nil, errors.New("log level not found")
	}

	return &loggerConfig{
		logLevel: logLevel,
	}, nil
}

func (c *config) LogLevel() string {
	return c.loglevel
}
