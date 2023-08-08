package configparser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RunConfig struct {
		Domains        []string `yaml:"domains"`
		SlackWebhook   string   `yaml:"slack"`
		DiscordWebhook string   `yaml:"discord"`
		RunType        string   `yaml:"runType"`
		ActiveWordlist string   `yaml:"activeWordList"`
		ActiveThreads  int      `yaml:"activeThreads"`
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
