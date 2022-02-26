package configs

type Server struct {
	Port int
	LogLevel string
}

type Configuration struct {
	Server `yaml:"server"`
}