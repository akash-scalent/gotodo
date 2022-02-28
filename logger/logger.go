package logger

import (
	"os"
	"strings"

	"github.com/akash-scalent/gotodo/configs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)


func InitializeLogger() *zerolog.Logger {
	logFile, err := os.OpenFile("todo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Str("service", "logger").Msg("Error opening log file")
	}

	multi := zerolog.MultiLevelWriter(logFile, os.Stdout)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Logger started")
	

	// Set logger level from configuration file

	if loggerLevel, err := zerolog.ParseLevel(strings.ToLower(configs.Config.LogLevel)); err != nil {
		logger = logger.Level(zerolog.InfoLevel)
		logger.Error().Str("service","logger").Err(err).Msg("Error parsing config level")
	} else {
		logger = logger.Level(loggerLevel)
	}
	logger.Info().Str("service", "configuration").Interface("config", configs.Config).Msg("Configuration")
	return &logger
}