package configparser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RunConfig struct {
		Domains      []string `yaml:"domains"`
		SlackToken   string   `yaml:"slack"`
		DiscordToken string   `yaml:"discord"`
	} `yaml:"runConfig"`
}

func ParseConfig() Config {
	// Read file data
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Initialize configuration
	var config Config

	// Unmarshal YAML data into Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Print out the parsed data
	fmt.Printf("Parsed config: %+v\n", config)
	return config
}
