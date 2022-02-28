package configs

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Server struct {
	SPort int `yaml:"port"`
	LogLevel string
}
type Database struct {
	Host string
  User string
  Password string
  Dbname string
  DPort int `yaml:"port"`
  Timezone string
}
type Configuration struct {
	Server `yaml:"server"`
	Database `yaml:"database"`
}

var Config Configuration

func InitializeConfiguration() {
	configFile, err := os.Open(filepath.Join("configs", "config.yaml"))
	if err != nil {
		log.Fatal().Err(err).Str("service", "configuration").Msg("Error reading configurating file")
	}
	err = yaml.NewDecoder(configFile).Decode(&Config)
	configFile.Close()
	if err != nil {
		log.Fatal().Err(err).Str("service", "configuration").Msg("Error unmarshalling YAML file")
	}
}