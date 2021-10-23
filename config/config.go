package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Default struct {
	LogLevel string
}

type Server struct {
	Host string
	Port string
}

type Postgres struct {
	Dev      bool
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

type Config struct {
	Default  Default
	Server   Server
	Postgres Postgres
}

// Return configuration
func GetConfig(path string) *Config {
	confFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf Config

	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}
