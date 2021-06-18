package main

import (
	"URLShortener/internal/URLShortener"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// parseConfigPath returns path of the config file
func parseConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config-path", "config/config.yml", "path to configuration file")
	flag.Parse()
	return configPath
}

// parseConfigFile returns config struct from .yaml file
func parseConfigFile(configPath string) *URLShortener.Config {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	c := URLShortener.NewConfig()
	if err := yaml.Unmarshal(file, c); err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}
	return c
}

// main start server
func main() {
	configPath := parseConfigPath()
	config := parseConfigFile(configPath)
	rand.Seed(time.Now().UnixNano())
	srv := URLShortener.NewGRPCServer(config)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
