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

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yml", "path to configuration file")
	rand.Seed(time.Now().UnixNano())
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

// main configure and start URLShortener server
func main() {
	flag.Parse()
	config := parseConfigFile(configPath)

	srv := URLShortener.NewGRPCServer(config)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
