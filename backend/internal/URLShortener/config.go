package URLShortener

import "URLShortener/internal/store"

type Config struct {
	Server *ServerConfig
	Store  *store.Config
}

type ServerConfig struct {
	BindAddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		BindAddr: ":8080",
		LogLevel: "info",
	}
}

func NewConfig() *Config {
	return &Config{
		Server: NewServerConfig(),
		Store:  store.NewConfig(),
	}
}
