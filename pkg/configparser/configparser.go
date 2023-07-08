package configparser

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RunConfig struct {
		Seed_domain string `yaml:"seedDomain"`
	} `yaml:"runConfig"`
	// Database struct {
	//     Username string `yaml:"user"`
	//     Password string `yaml:"pass"`
	// } `yaml:"database"`
}

func ParseConfig() Config {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("Failed to load config file!")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Failed to Decode/parse config!")
	}
	return cfg
}
