package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/akash-scalent/gotodo/configs"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"

	"github.com/rs/zerolog"
)

var Configuration = &configs.Configuration{}

func main() {
	logFile, err := os.OpenFile("todo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Str("service", "logger").Msg("Error opening log file")
	}

	multi := zerolog.MultiLevelWriter(logFile, os.Stdout)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Logger started")
	configFile, err := os.Open(filepath.Join("configs", "config.yaml"))
	if err != nil {
		logger.Fatal().Err(err).Str("service", "configuration").Msg("Error reading configurating file")
	}
	err = yaml.NewDecoder(configFile).Decode(Configuration)
	configFile.Close()
	if err != nil {
		logger.Fatal().Err(err).Str("service", "configuration").Msg("Error unmarshalling YAML file")
	}

	// Set logger level from configuration file

	if loggerLevel, err := zerolog.ParseLevel(strings.ToLower(Configuration.LogLevel)); err != nil {
		logger = logger.Level(zerolog.InfoLevel)
	} else {
		logger.Error().Err(err).Send()
		logger = logger.Level(loggerLevel)
	}
	logger.Info().Str("service", "configuration").Interface("config", Configuration).Msg("Configuration")
}
