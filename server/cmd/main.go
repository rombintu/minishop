package main

import (
	"flag"
	"log"

	"github.com/rombintu/minishop/config"
	"github.com/rombintu/minishop/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
}

// My code is comments
func main() {
	flag.Parse()
	config := config.GetConfig(configPath)

	s := app.NewApp(config)
	if err := s.Start(); err != nil {
		log.Fatalf("%v", err)
	}
}
